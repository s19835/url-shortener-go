package handlers

import (
	"net/http"

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

	shortCode, err := h.service.ShortenURL(c.Request.Context(), req.URL, *req.ExpiresIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// prepare the short url
	shortURL := "http://" + c.Request.Host + "/" + shortCode
	c.JSON(http.StatusCreated, models.ShortenResponse{ShortURL: shortURL})
}
