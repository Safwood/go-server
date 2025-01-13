package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authenticationsHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authenticationsHeader)

	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "Empty token")
		return
	}
	
	headerPart := strings.Split(header, " ")
	
	if len(headerPart) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "Not valid token")
		return
	}
	
	userId, err := h.services.ParseToken(headerPart[1])
    if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "Not valid token")
		return
	}

	c.Set(userCtx, userId)
}