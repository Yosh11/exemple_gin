package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Error ...
type error struct {
	Message string `json:"message"`
}

// NewErrorResposnse handler errors
func NewErrorResposnse(c *gin.Context, statusCode int, message string) {
	log.Error(message)

	c.AbortWithStatusJSON(statusCode, error{message})
}
