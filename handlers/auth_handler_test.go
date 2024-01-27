package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	url := "/api/v1/login"
	r := gin.Default()
	authHandler := NewAuthHandler()
	r.POST(url, authHandler.Login)
	creds := Credentials{
		Email:    "admin@admin.com",
		Password: "password123",
	}
	payload, _ := json.Marshal(creds)
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
