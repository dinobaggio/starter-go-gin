package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"starter-go-gin/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	url := "/api/v1/auth/login"
	r := gin.Default()
	authHandler := NewAuthHandler()
	r.POST(url, authHandler.Login)
	creds := Credentials{
		Email:    "admin@admin.com",
		Password: "password123",
	}
	payload, err := json.Marshal(creds)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	responseData, err := io.ReadAll(w.Body)

	if err != nil {
		t.Fatal(err)
	}

	var response map[string]interface{}

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		t.Fatal(err)
	}

	data := response["data"].(map[string]interface{})
	user := data["user"].(map[string]interface{})

	assert.Equal(t, creds.Email, user["email"])
	assert.Equal(t, "string", reflect.TypeOf(data["token"]).String())
}

func TestRegrister(t *testing.T) {
	url := "/api/v1/auth/register"
	r := gin.Default()
	authHandler := NewAuthHandler()
	r.POST(url, authHandler.Register)
	formUser := FormRegister{
		Name:            "test",
		Email:           "test@test.com",
		Password:        "password123",
		ConfirmPassword: "password123",
	}
	form, err := json.Marshal(formUser)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(form))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, err := io.ReadAll(w.Body)

	if err != nil {
		t.Fatal(err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(responseData, &response)

	if err != nil {
		t.Fatal(err)
	}

	data := response["data"].(map[string]interface{})
	user := data["user"].(map[string]interface{})

	assert.Equal(t, formUser.Email, user["email"])
	assert.Equal(t, formUser.Name, user["name"])

	newUser := models.NewUser()
	err = newUser.DeleteByEmail(formUser.Email)

	if err != nil {
		t.Fatal(err)
	}
}
