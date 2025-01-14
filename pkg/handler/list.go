package handler

import (
	"net/http"

	todo "github.com/Safwood/go-server"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	listId, err := h.services.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"listId": listId,
	})
}

func (h *Handler) getAllLists(c *gin.Context)  {
	
}

func (h *Handler) getListById(c *gin.Context)  {
	
}

func (h *Handler) updateList(c *gin.Context)  {
	
}

func (h *Handler) deleteList(c *gin.Context)  {
	
}