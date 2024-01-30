package handlers

import (
	"fmt"
	"starter-go-gin/libs"
	"starter-go-gin/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (*UserHandler) List(c *gin.Context) {
	users, err := models.NewUser().GetUsersWithPagination()

	if err != nil {
		fmt.Println(err.Error())
		libs.ResponseInternalServerError(c)
		return
	}

	libs.ResponseSuccessWithData(c, users)
}
