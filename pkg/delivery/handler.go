package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/tetovske/advertisement-service/pkg/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(serv *services.Service) *Handler {
	return &Handler{services: serv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		ads := api.Group("/advertisements")
		{
			ads.GET("/", h.getAdvertisementsList)
			ads.GET("/:id", h.getAdvertisement)
			ads.POST("/", h.createAdvertisement)
		}
	}

	return router
}