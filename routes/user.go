package routes

import (
	"starter-go-gin/handlers"
	"starter-go-gin/libs"

	"github.com/gin-gonic/gin"
)

func UserRouter(group *gin.RouterGroup) {
	userHandler := handlers.NewUserHandler()

	router := group.Group("users")
	router.Use(libs.ValidToken())
	{
		router.GET("", userHandler.List)
	}
}
