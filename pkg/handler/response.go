package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Error ...
type errorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse handler errors
func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Error(message)

	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
