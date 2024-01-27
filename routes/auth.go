package routes

import (
	"starter-go-gin/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(group *gin.RouterGroup) {
	authHandler := handlers.NewAuthHandler()
	router := group.Group("auth")
	{
		router.POST("/login", authHandler.Login)
		router.POST("/register", authHandler.Register)
	}
}
