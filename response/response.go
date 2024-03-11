package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	success = iota
	fail
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    success,
		"message": "success",
	})
}

func SuccessWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    success,
		"message": message,
	})
}

func SuccessWithData(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    success,
		"data":    data,
		"message": message,
	})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    fail,
		"message": "error",
	})
}

func FailWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    fail,
		"message": message,
	})
}

func FailWithData(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    fail,
		"data":    data,
		"message": message,
	})
}
