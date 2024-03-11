package example

import (
	"awesomeProject/api"

	"github.com/gin-gonic/gin"
)

type ExampleRouter struct {
}

func (*ExampleRouter) InitExample(r *gin.Engine) {
	group := r.Group("/example")
	apiGroup := api.ApiGroupApp.ExampleApiGroup
	group.GET("/ping", apiGroup.Test)
}
