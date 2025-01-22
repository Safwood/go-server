package handler

import (
    "github.com/Safwood/go-server/pkg/service"
	"github.com/gin-gonic/gin"

    "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

    _ "github.com/Safwood/go-server/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
    return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
            }
        }

        items := api.Group("items")
        {
            items.GET("/:id", h.getItemById)
            items.PUT("/:id", h.updateItem)
            items.DELETE("/:id", h.deleteItem)
        }

        parks := api.Group("/parks", h.UserIdentity) 
        {
            parks.POST("/", h.createPark)
            parks.GET("/", h.getAllParks)
            parks.GET("/:id", h.getParkById)
            parks.PUT("/:id", h.updatePark)
            parks.DELETE("/:id", h.deletePark)
        }
    }


    return router
}