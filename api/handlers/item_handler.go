package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/api/dto"
	"github.com/ParthPoudyal/E-commerce-project/api/middleware"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/bid"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/user"
)

type ItemHandler struct {
	itemRepo item.ItemRepo
	bidRepo  bid.BidRepository
	userRepo user.UserRepo
}

func NewItemHandler(iR item.ItemRepo, bR bid.BidRepository, uR user.UserRepo) *ItemHandler {
	return &ItemHandler{
		itemRepo: iR,
		bidRepo:  bR,
		userRepo: uR,
	}
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	u, err := h.userRepo.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.APIResponse{Success: false, Error: "User not found"})
		return
	}

	var req dto.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	now := time.Now()
	status := domain.AVAILABLE
	var validTo *time.Time

	if req.IsAuction {
		status = domain.BIDDING
		days := req.DurationDays
		if days <= 0 {
			days = 7 // Default 7-day auction
		}
		endTime := now.Add(time.Duration(days) * 24 * time.Hour)
		validTo = &endTime
		if req.StartingBid != nil {
			req.Price = *req.StartingBid
		}
	}

	itemObj := &domain.Item{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		ImageURL:    req.ImageURL,
		PostedBy:    u,
		Price:       req.Price,
		Status:      status,
		ValidFrom:   now,
		ValidTo:     validTo,
		IsAuction:   req.IsAuction,
		StartingBid: req.StartingBid,
		Base: domain.Base{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	if err := h.itemRepo.AddItem(itemObj); err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Data:    itemObj,
	})
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: "Invalid item ID"})
		return
	}

	itm, err := h.itemRepo.FindItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.APIResponse{Success: false, Error: "Item not found"})
		return
	}

	highestBid, _ := h.bidRepo.GetHighestBidForItem(id)
	bids, _ := h.bidRepo.GetBidHistoryForItem(id)

	var timeLeftSec int64
	if itm.ValidTo != nil {
		diff := time.Until(*itm.ValidTo).Seconds()
		if diff > 0 {
			timeLeftSec = int64(diff)
		}
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data: dto.ItemResponse{
			Item:        *itm,
			HighestBid:  highestBid,
			TotalBids:   len(bids),
			TimeLeftSec: timeLeftSec,
		},
	})
}

func (h *ItemHandler) ListItems(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	categoryIDStr := c.Query("category_id")
	isAuctionQuery := c.Query("is_auction")

	var items []domain.Item
	var err error

	if categoryIDStr != "" {
		if catID, parseErr := uuid.Parse(categoryIDStr); parseErr == nil {
			items, err = h.itemRepo.FindItemsByCategory(catID, limit, offset)
		}
	} else if isAuctionQuery == "true" {
		items, err = h.itemRepo.FindAuctionItems(limit, offset)
	} else {
		items, err = h.itemRepo.ListAllItems(limit, offset)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    items,
	})
}

func (h *ItemHandler) ListAuctions(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	items, err := h.itemRepo.FindAuctionItems(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    items,
	})
}

func (h *ItemHandler) SearchItems(c *gin.Context) {
	query := c.Query("q")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	items, err := h.itemRepo.SearchItemByUser(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    items,
	})
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: "Invalid item ID"})
		return
	}

	itm, err := h.itemRepo.FindItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.APIResponse{Success: false, Error: "Item not found"})
		return
	}

	if itm.PostedBy != nil && itm.PostedBy.ID != userID {
		c.JSON(http.StatusForbidden, dto.APIResponse{Success: false, Error: "You can only update your own items"})
		return
	}

	var req dto.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	if req.Name != "" {
		itm.Name = req.Name
	}
	if req.Description != "" {
		itm.Description = req.Description
	}
	if req.CategoryID != nil {
		itm.CategoryID = req.CategoryID
	}
	if req.ImageURL != "" {
		itm.ImageURL = req.ImageURL
	}
	if req.Price != nil {
		itm.Price = *req.Price
	}
	if req.Status != nil {
		itm.Status = *req.Status
	}

	if err := h.itemRepo.UpdateItem(itm); err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    itm,
	})
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: "Invalid item ID"})
		return
	}

	itm, err := h.itemRepo.FindItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.APIResponse{Success: false, Error: "Item not found"})
		return
	}

	if itm.PostedBy != nil && itm.PostedBy.ID != userID {
		c.JSON(http.StatusForbidden, dto.APIResponse{Success: false, Error: "You can only delete your own items"})
		return
	}

	if err := h.itemRepo.DeleteItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Item deleted successfully",
	})
}

func (h *ItemHandler) GetItemsByUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: "Invalid user ID"})
		return
	}

	items, err := h.itemRepo.FindItemByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    items,
	})
}
