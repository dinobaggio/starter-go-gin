package handlers

import (
	"starter-go-gin/libs"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (*AuthHandler) Login(c *gin.Context) {
	libs.ResponseSuccessWithData(c, gin.H{})
}
