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
	req, err := http.NewRequest("GET", url, nil)
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

	fmt.Println(string(responseData))
}
