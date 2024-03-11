package initiallize

import (
	"awesomeProject/middleware"
	"awesomeProject/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors)
	exampleGroup := router.RouterGroupApp.ExampleRouterGroup
	exampleGroup.InitExample(r)
	kubectlGroup := router.RouterGroupApp.KubectlRouterGroup
	kubectlGroup.InitK8S(r)
	return r
}
