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

func (r *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		ads := api.Group("/advertisements")
		{
			ads.GET("/", r.getAdvertisementsList)
			ads.GET("/:id", r.getAdvertisement)
			ads.POST("/", r.createAdvertisement)
		}
	}

	return router
}