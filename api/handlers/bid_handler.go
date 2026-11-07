package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/api/dto"
	"github.com/ParthPoudyal/E-commerce-project/api/middleware"
	"github.com/ParthPoudyal/E-commerce-project/api/ws"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/bid"
	bidservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/bid/services"
)

type BidHandler struct {
	placeBidService *bidservices.PlaceBidService
	bidRepo         bid.BidRepository
	wsHub           *ws.Hub
}

func NewBidHandler(
	placeBidService *bidservices.PlaceBidService,
	bR bid.BidRepository,
	hub *ws.Hub,
) *BidHandler {
	return &BidHandler{
		placeBidService: placeBidService,
		bidRepo:         bR,
		wsHub:           hub,
	}
}

func (h *BidHandler) PlaceBid(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	var req dto.PlaceBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	result, err := h.placeBidService.Execute(req.ItemID, userID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	// Broadcast live bid to all participants in this auction's WebSocket room
	if h.wsHub != nil {
		h.wsHub.BroadcastBid(req.ItemID, result.Bid)

		// If a previous highest bidder was outbid, notify them
		if result.PreviousHighestUser != nil && *result.PreviousHighestUser != userID {
			h.wsHub.BroadcastOutbid(req.ItemID, *result.PreviousHighestUser, req.Amount.StringFixed(2))
		}
	}

	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Data: dto.BidResponse{
			Bid:                 result.Bid,
			PreviousHighestUser: result.PreviousHighestUser,
			IsNewHighest:        result.IsNewHighest,
		},
	})
}

func (h *BidHandler) GetBidHistory(c *gin.Context) {
	itemIDStr := c.Param("id")
	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: "Invalid item ID"})
		return
	}

	bids, err := h.bidRepo.GetBidHistoryForItem(itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    bids,
	})
}

func (h *BidHandler) GetUserBids(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	bids, err := h.bidRepo.GetBidsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    bids,
	})
}

