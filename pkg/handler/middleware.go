package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const userCtx = "userID"

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewErrorResposnse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	// valid Authorization header have 2 string bearer and token
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResposnse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResposnse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}
