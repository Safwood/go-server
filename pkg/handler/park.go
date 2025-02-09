package handler

import (
	"net/http"
	"strconv"

	sights "github.com/Safwood/go-server"
	"github.com/gin-gonic/gin"
)

// @Summary Create park
// @Security ApiKeyAuth
// @Tags parks
// @Description create park
// @ID create-park
// @Accept  json
// @Produce  json
// @Param input body sights.Park true "park info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/parks [post]
func (h *Handler) createPark(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

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

// @Summary GetAllParks 
// @Security ApiKeyAuth
// @Tags parks
// @Description get list of parks
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} []sights.Park
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/parks [get]
func (h *Handler) getAllParks(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
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

// @Summary Get Park By Id
// @Security ApiKeyAuth
// @Tags parks
// @Description get park by id
// @ID get-park-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} sights.Park
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/parks/:id [get]
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

// @Summary Update Park
// @Security ApiKeyAuth
// @Tags parks
// @Description update park
// @ID update-park
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/parks/:id [put]
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

// @Summary Delete Park
// @Security ApiKeyAuth
// @Tags parks
// @Description delete park
// @ID delete-park
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/parks/:id [delete]
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

