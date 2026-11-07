package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/api/dto"
)

type UploadHandler struct {
	uploadDir string
}

func NewUploadHandler(uploadDir string) *UploadHandler {
	_ = os.MkdirAll(uploadDir, 0755)
	return &UploadHandler{
		uploadDir: uploadDir,
	}
}

func (h *UploadHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "No image file provided in form data ('image')",
		})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" && ext != ".gif" {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Unsupported file type. Only JPEG, PNG, WEBP, and GIF are allowed.",
		})
		return
	}

	newFileName := fmt.Sprintf("%s_%d%s", uuid.New().String()[:12], time.Now().Unix(), ext)
	dst := filepath.Join(h.uploadDir, newFileName)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to save uploaded image: %v", err),
		})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	fileURL := fmt.Sprintf("%s://%s/uploads/%s", scheme, c.Request.Host, newFileName)

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data: gin.H{
			"url":      fileURL,
			"filename": newFileName,
		},
	})
}

