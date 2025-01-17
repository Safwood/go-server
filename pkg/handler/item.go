package handler

import (
	"net/http"
	"strconv"

	todo "github.com/Safwood/go-server"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id")); 
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	itemId, err := h.services.TodoItem.CreateItem(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"itemId": itemId,
	})
}

func (h *Handler) getAllItems(c *gin.Context)  {
	
}

func (h *Handler) getItemById(c *gin.Context)  {
	
}

func (h *Handler) updateItem(c *gin.Context)  {
	
}

func (h *Handler) deleteItem(c *gin.Context)  {
	
}