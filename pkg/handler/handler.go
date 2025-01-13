package handler

import (
	"github.com/Safwood/go-server/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
    return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

    auth := router.Group("/auth") 
    {
          auth.POST("/sign-up", h.signUp)
          auth.POST("/sign-in", h.signIn)
    }

    api := router.Group("/api", h.UserIdentity) 
    {
        lists := api.Group("/lists")
        {
            lists.POST("/", h.createList)
            lists.GET("/", h.getAllLists)
            lists.GET("/:id", h.getListById)
            lists.PUT("/:id", h.updateList)
            lists.DELETE("/:id", h.deleteList)

            item := lists.Group(":id/items")
            {
                item.POST("/", h.createItem)
                item.GET("/", h.getAllItems)
                item.GET("/:item_id", h.getItemById)
                item.PUT("/:item_id", h.updateItem)
                item.DELETE("/:item_id", h.deleteItem)
            }
        }
    }

    return router
}