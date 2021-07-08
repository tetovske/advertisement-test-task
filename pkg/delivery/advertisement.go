package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Handler) createAdvertisement(c *gin.Context) {

}

func (r *Handler) getAdvertisement(c *gin.Context) {

}

func (r *Handler) getAdvertisementsList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": 1,
	})
}
