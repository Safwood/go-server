package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/safwood/go-server/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/safwood/go-server/docs"
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