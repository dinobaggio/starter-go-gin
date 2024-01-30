package handlers

import (
	"starter-go-gin/libs"
	"starter-go-gin/models"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (*AuthHandler) Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		libs.ResponseBadRequest(c, err)
		return
	}

	user, err := models.LoginCheck(creds.Email, creds.Password)
	if err != nil {
		libs.ResponseBadRequest(c, err)
		return
	}

	token, err := libs.GenerateToken(user.ID, user.Email)

	if err != nil {
		libs.ResponseBadRequest(c, err)
		return
	}

	libs.ResponseSuccessWithData(c, gin.H{
		"user":  &user,
		"token": token,
	})
}

type FormRegister struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Name            string `json:"name"`
}

func (*AuthHandler) Register(c *gin.Context) {
	var formRegister FormRegister

	if err := c.ShouldBindJSON(&formRegister); err != nil {
		libs.ResponseBadRequest(c, err)
		return
	}

	user := models.NewUser()
	user.Name = formRegister.Name
	user.Email = formRegister.Email

	hashPass, err := models.GeneratePassword(formRegister.Password)
	if err != nil {
		libs.ResponseBadRequest(c, err)
		return
	}

	user.Password = hashPass

	lastID, err := user.Save()
	if err != nil {
		libs.ResponseBadRequest(c, err)
		return
	}

	userData, err := user.GetByID(lastID)
	if err != nil {
		libs.ResponseBadRequest(c, err)
		return
	}

	libs.ResponseSuccessWithData(c, gin.H{
		"user": userData,
	})
}
