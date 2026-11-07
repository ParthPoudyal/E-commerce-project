package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ParthPoudyal/E-commerce-project/api/dto"
	"github.com/ParthPoudyal/E-commerce-project/api/middleware"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/payment"
	paymentservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/payment/services"
)

type PaymentHandler struct {
	khaltiService *paymentservices.KhaltiService
	paymentRepo   payment.PaymentRepo
}

func NewPaymentHandler(kSvc *paymentservices.KhaltiService, pR payment.PaymentRepo) *PaymentHandler {
	return &PaymentHandler{
		khaltiService: kSvc,
		paymentRepo:   pR,
	}
}

func (h *PaymentHandler) InitiateKhaltiPayment(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	var req dto.InitiatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	out, err := h.khaltiService.InitiatePayment(paymentservices.InitiatePaymentInput{
		ItemID: req.ItemID,
		UserID: userID,
		Amount: req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    out,
	})
}

func (h *PaymentHandler) VerifyKhaltiPayment(c *gin.Context) {
	var req dto.VerifyPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	tx, err := h.khaltiService.VerifyPayment(req.PIDX)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    tx,
	})
}

func (h *PaymentHandler) GetUserPayments(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Success: false, Error: "Unauthorized"})
		return
	}

	payments, err := h.paymentRepo.ListPaymentsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    payments,
	})
}

