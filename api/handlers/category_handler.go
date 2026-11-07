package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ParthPoudyal/E-commerce-project/api/dto"
	categoryservices "github.com/ParthPoudyal/E-commerce-project/internal/usecase/category/services"
)

type CategoryHandler struct {
	categoryService *categoryservices.CategoryService
}

func NewCategoryHandler(cSvc *categoryservices.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: cSvc,
	}
}

func (h *CategoryHandler) ListCategories(c *gin.Context) {
	categories, err := h.categoryService.ListCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    categories,
	})
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
	slug := c.Param("slug")
	cat, err := h.categoryService.GetCategoryBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.APIResponse{Success: false, Error: "Category not found"})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    cat,
	})
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	cat, err := h.categoryService.CreateCategory(req.Name, req.Description, req.ImageURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Data:    cat,
	})
}

