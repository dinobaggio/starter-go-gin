package routes

import (
	"starter-go-gin/handlers"

	"github.com/gin-gonic/gin"
)

func UserRouter(group *gin.RouterGroup) {
	userHandler := handlers.NewUserHandler()
	router := group.Group("users")
	{
		router.GET("", userHandler.List)
	}
}
