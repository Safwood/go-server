package handler

import (
	"errors"
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
		NewErrorResponse(c, http.StatusUnauthorized, "Parse token error")
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "User id is not found")
		return 0, errors.New("user id not found")
	}
	
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "user id is not of valid type")
		return 0, errors.New("user id is not of valid type")
	}

	return idInt, nil
}