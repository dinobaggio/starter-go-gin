package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListUser(t *testing.T) {
	r := gin.Default()
	userHandler := NewUserHandler()
	url := "/api/v1/users"
	r.GET(url, userHandler.List)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	fmt.Println(string(responseData))
}
