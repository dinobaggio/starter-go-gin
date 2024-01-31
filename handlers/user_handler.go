package handlers

import (
	"starter-go-gin/libs"
	"starter-go-gin/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type Query struct {
	Page  int `form:"page,default=1" json:"page"`
	Limit int `form:"limit,default=10" json:"limit"`
}

func (*UserHandler) List(c *gin.Context) {
	var query Query

	if err := c.BindQuery(&query); err != nil {
		libs.ResponseInternalServerError(c)
		return
	}

	users, err := models.NewUser().GetUsersWithPagination(
		int64(query.Page),
		int64(query.Limit),
	)

	if err != nil {
		libs.ResponseInternalServerError(c)
		return
	}

	libs.ResponseSuccessWithData(c, gin.H{
		"users": users,
		"query": query,
	})
}
