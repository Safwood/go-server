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