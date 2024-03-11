package kubectl

import (
	"awesomeProject/api"

	"github.com/gin-gonic/gin"
)

type KubectlRouter struct {
}

func (*KubectlRouter) InitK8S(r *gin.Engine) {
	group := r.Group("/kubectl")
	apiGroup := api.ApiGroupApp.KubectlApiGroup
	group.GET("/getPodsNum", apiGroup.GetPodsNumApi)
}
