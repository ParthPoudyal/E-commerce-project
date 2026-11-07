package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/auth"
)

type BroadcastMessage struct {
	AuctionID uuid.UUID
	Payload   []byte
}

type Hub struct {
	// Registered clients mapped by auction ID
	auctionRooms map[uuid.UUID]map[*Client]bool

	// Inbound messages from the clients / services
	Broadcast chan *BroadcastMessage

	// Register requests from the clients
	Register chan *Client

	// Unregister requests from clients
	Unregister chan *Client

	mu sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:    make(chan *BroadcastMessage, 256),
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		auctionRooms: make(map[uuid.UUID]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			if _, ok := h.auctionRooms[client.AuctionID]; !ok {
				h.auctionRooms[client.AuctionID] = make(map[*Client]bool)
			}
			h.auctionRooms[client.AuctionID][client] = true
			roomSize := len(h.auctionRooms[client.AuctionID])
			h.mu.Unlock()

			// Send welcome message with viewer count
			welcome, _ := json.Marshal(map[string]interface{}{
				"type":        "room_joined",
				"auction_id":  client.AuctionID,
				"viewer_count": roomSize,
				"timestamp":   time.Now(),
			})
			select {
			case client.Send <- welcome:
			default:
			}

		case client := <-h.Unregister:
			h.mu.Lock()
			if clients, ok := h.auctionRooms[client.AuctionID]; ok {
				if _, exists := clients[client]; exists {
					delete(clients, client)
					close(client.Send)
					if len(clients) == 0 {
						delete(h.auctionRooms, client.AuctionID)
					}
				}
			}
			h.mu.Unlock()

		case message := <-h.Broadcast:
			h.mu.RLock()
			clients, ok := h.auctionRooms[message.AuctionID]
			if ok {
				for client := range clients {
					select {
					case client.Send <- message.Payload:
					default:
						close(client.Send)
						delete(clients, client)
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) BroadcastBid(auctionID uuid.UUID, bidData interface{}) {
	payload, err := json.Marshal(map[string]interface{}{
		"type":      "bid_placed",
		"data":      bidData,
		"timestamp": time.Now(),
	})
	if err == nil {
		h.Broadcast <- &BroadcastMessage{
			AuctionID: auctionID,
			Payload:   payload,
		}
	}
}

func (h *Hub) BroadcastOutbid(auctionID, previousBidderID uuid.UUID, newAmount string) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	clients, ok := h.auctionRooms[auctionID]
	if !ok {
		return
	}

	payload, _ := json.Marshal(map[string]interface{}{
		"type":        "outbid",
		"auction_id":  auctionID,
		"message":     "You have been outbid!",
		"new_highest": newAmount,
		"timestamp":   time.Now(),
	})

	for client := range clients {
		if client.UserID == previousBidderID {
			select {
			case client.Send <- payload:
			default:
			}
		}
	}
}

func (h *Hub) HandleWebSocket(c *gin.Context) {
	auctionIDStr := c.Param("id")
	auctionID, err := uuid.Parse(auctionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid auction ID"})
		return
	}

	// Optional token from query param ?token=...
	token := c.Query("token")
	var userID uuid.UUID
	if token != "" {
		if claims, err := auth.ValidateToken(token); err == nil {
			userID = claims.UserID
		}
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade websocket: %v", err)
		return
	}

	client := &Client{
		Hub:       h,
		Conn:      conn,
		Send:      make(chan []byte, 256),
		AuctionID: auctionID,
		UserID:    userID,
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}

