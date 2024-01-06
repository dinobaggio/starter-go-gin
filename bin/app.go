package bin

import (
	"fmt"
	"log"
	"net/http"
	"starter-go-gin/routes"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func (a *App) Initialize() {
	a.Router = gin.Default()
	a.setRoutes()
}

// setRoutes sets up the routes for the application.
func (a *App) setRoutes() {
	router := a.Router
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/healthcheck", func(c *gin.Context) {
				c.JSON(http.StatusOK, "It's work")
			})
			routes.UserRouter(v1)
		}
	}
}

// Run starts the application.
func (a *App) Run(addr string) {
	fmt.Println("running", addr)
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.Panic("Failed to start the server:", err)
	}
}

func NewApp() *App {
	app := App{}
	app.Initialize()

	return &app
}
