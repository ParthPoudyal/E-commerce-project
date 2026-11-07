package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/api/dto"
	"github.com/ParthPoudyal/E-commerce-project/api/middleware"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
	wishlistservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/wishlist/services"
)

type WishlistHandler struct {
	wishlistService *wishlistservices.WishlistService
	itemRepo        item.ItemRepo
}

func NewWishlistHandler(wSvc *wishlistservices.WishlistService, iR item.ItemRepo) *WishlistHandler {
	return &WishlistHandler{
		wishlistService: wSvc,
		itemRepo:        iR,
	}
}

func (h *WishlistHandler) GetWishlist(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	items, err := h.wishlistService.GetUserWishlist(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	hydrated := make([]dto.WishlistResponseItem, len(items))
	for i, wi := range items {
		itm, _ := h.itemRepo.FindItemByID(wi.ItemId)
		hydrated[i] = dto.WishlistResponseItem{
			WishlistItem: wi,
			Item:         itm,
		}
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    hydrated,
	})
}

func (h *WishlistHandler) AddToWishlist(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	var req dto.AddToWishlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	wi, err := h.wishlistService.AddToWishlist(userID, req.ItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Data:    wi,
	})
}

func (h *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	itemIDStr := c.Param("itemId")
	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: "Invalid item ID"})
		return
	}

	if err := h.wishlistService.RemoveFromWishlist(userID, itemID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Item removed from wishlist",
	})
}

func (h *WishlistHandler) CheckWishlisted(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusOK, dto.APIResponse{Success: true, Data: false})
		return
	}

	itemIDStr := c.Param("itemId")
	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: "Invalid item ID"})
		return
	}

	isWishlisted, _ := h.wishlistService.IsWishlisted(userID, itemID)
	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    isWishlisted,
	})
}

