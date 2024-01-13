package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "internal server error",
	})
}

func ResponseSuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}
