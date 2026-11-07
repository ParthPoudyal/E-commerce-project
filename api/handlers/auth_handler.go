package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ParthPoudyal/E-commerce-project/api/dto"
	"github.com/ParthPoudyal/E-commerce-project/api/middleware"
	authservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/auth/services"
)

type AuthHandler struct {
	authService *authservices.AuthService
}

func NewAuthHandler(authService *authservices.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	result, err := h.authService.Register(authservices.RegisterInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		PhoneNo:  req.PhoneNo,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Data:    result,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	result, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    result,
	})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error:   "Unauthorized",
		})
		return
	}

	u, err := h.authService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.APIResponse{
			Success: false,
			Error:   "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    u,
	})
}

