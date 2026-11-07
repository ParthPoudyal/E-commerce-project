package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/api/dto"
	"github.com/ParthPoudyal/E-commerce-project/api/middleware"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/item"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/purchase"
	purchaseservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/purchase/services"
)

type HydratedPurchase struct {
	Purchase domain.Purchase `json:"purchase"`
	Item     *domain.Item    `json:"item"`
}

type PurchaseHandler struct {
	purchaseRepo  purchase.PurchaseRepo
	itemRepo      item.ItemRepo
	buyNowService *purchaseservices.BuyNowService
}

func NewPurchaseHandler(
	pR purchase.PurchaseRepo,
	iR item.ItemRepo,
	buyNowSvc *purchaseservices.BuyNowService,
) *PurchaseHandler {
	return &PurchaseHandler{
		purchaseRepo:  pR,
		itemRepo:      iR,
		buyNowService: buyNowSvc,
	}
}

func (h *PurchaseHandler) GetPurchases(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	purchases, err := h.purchaseRepo.ListPurchasesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	hydrated := make([]HydratedPurchase, len(purchases))
	for i, p := range purchases {
		itm, _ := h.itemRepo.FindItemByID(p.ItemID)
		hydrated[i] = HydratedPurchase{
			Purchase: p,
			Item:     itm,
		}
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    hydrated,
	})
}

func (h *PurchaseHandler) BuyNow(c *gin.Context) {
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

	itm, err := h.itemRepo.FindItemByID(itemID)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.APIResponse{Success: false, Error: "Item not found"})
		return
	}

	purch, err := h.buyNowService.Execute(itemID, userID, itm.Price.InexactFloat64())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Data:    purch,
	})
}

