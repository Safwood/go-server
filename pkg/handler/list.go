package handler

import (
	"net/http"
	"strconv"

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
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	listId, err := h.services.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"listId": listId,
	})
}

type allListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAllLists(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allListsResponse{Data: lists})
}


func (h *Handler) getListById(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id")); 
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetListById(userId,listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context)  {
	
}

func (h *Handler) deleteList(c *gin.Context)  {
	
}