package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(group *gin.RouterGroup) {
	router := group.Group("users")
	{
		router.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
	}
}
