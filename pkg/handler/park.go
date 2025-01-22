package handler

import (
	"net/http"
	"strconv"

	sights "github.com/Safwood/go-server"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPark(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	// logFile, err := os.OpenFile("log.txt", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// mw := io.MultiWriter(os.Stdout, logFile)
	// log.SetOutput(mw)

	var input sights.Park
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	id, err := h.services.Park.CreatePark(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type allParksResponse struct {
	Data []sights.Park `json:"data"`
}

func (h *Handler) getAllParks(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	
	parks, err := h.services.Park.GetAllParks(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allParksResponse{
		parks,
	})
}

type parkResponse struct {
	Data sights.GetParkOutput `json:"data"`
}

func (h *Handler) getParkById(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id")); 
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	
	item, err := h.services.Park.GetParkById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, parkResponse{
		item,
	})
}

func (h *Handler) updatePark(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input sights.UpdateParkInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Park.UpdatePark(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deletePark(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id")); 
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	
	if err := h.services.Park.DeletePark(userId, itemId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

