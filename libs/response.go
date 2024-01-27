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

func ResponseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "bad request",
		"error":   err.Error(),
	})
}

func ResponseSuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}
