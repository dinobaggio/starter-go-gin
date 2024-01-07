package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/healthcheck", func(c *gin.Context) {
				c.JSON(http.StatusOK, "It's work")
			})
			UserRouter(v1)
		}
	}
}
