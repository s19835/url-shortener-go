package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/s19835/url-shortener-go/internal/models"
	"github.com/s19835/url-shortener-go/internal/services"
)

type URLHandler struct {
	service *services.URLService
}

func NewURLHandler(service *services.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) ShortenURL(c *gin.Context) {
	var req models.ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// default expiry : 30 days if not specified
	if req.ExpiresIn == nil || *req.ExpiresIn <= 0 {
		defaultExpiry := 720 * time.Hour // 30 days
		req.ExpiresIn = &defaultExpiry
	}

	shortCode, err := h.service.ShortenURL(c.Request.Context(), req.URL, *req.ExpiresIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// prepare the short url
	shortURL := "http://" + c.Request.Host + "/" + shortCode
	c.JSON(http.StatusCreated, models.ShortenResponse{ShortURL: shortURL})
}

func (h *URLHandler) RedirectURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	originalURL, err := h.service.GetOriginalURL(c.Request.Context(), shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
