package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/api/dto"
	"github.com/ParthPoudyal/E-commerce-project/api/middleware"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/cart"
	cartservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/cart/services"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
)

type CartHandler struct {
	cartRepo         cart.CartRepo
	itemRepo         item.ItemRepo
	addToCartService *cartservices.AddtoCartService
	removeFromCartService *cartservices.RemoveFromCartService
}

func NewCartHandler(
	cR cart.CartRepo,
	iR item.ItemRepo,
	addSvc *cartservices.AddtoCartService,
	remSvc *cartservices.RemoveFromCartService,
) *CartHandler {
	return &CartHandler{
		cartRepo:              cR,
		itemRepo:              iR,
		addToCartService:      addSvc,
		removeFromCartService: remSvc,
	}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	cartItems, err := h.cartRepo.ListUserCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	hydrated := make([]dto.CartResponseItem, len(cartItems))
	for i, ci := range cartItems {
		itm, _ := h.itemRepo.FindItemByID(ci.ItemID)
		hydrated[i] = dto.CartResponseItem{
			CartItem: ci,
			Item:     itm,
		}
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    hydrated,
	})
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	var req dto.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	ci, err := h.addToCartService.Execute(userID, req.ItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Data:    ci,
	})
}

func (h *CartHandler) RemoveFromCart(c *gin.Context) {
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

	if err := h.removeFromCartService.Execute(userID, itemID); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Item removed from cart",
	})
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	if err := h.cartRepo.ClearCart(userID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Cart cleared",
	})
}

