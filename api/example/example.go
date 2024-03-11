package example

import (
	"awesomeProject/response"

	"github.com/gin-gonic/gin"
)

type ExampleApi struct {
}

func (*ExampleApi) Test(c *gin.Context) {
	response.SuccessWithMessage(c, "pong")
}
