package kubectl

import (
	"awesomeProject/global"
	"awesomeProject/response"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodsNumApi struct {
}

func (*PodsNumApi) GetPodsNumApi(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeClientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, item := range list.Items {
		fmt.Println(item.Namespace, item.Name)
	}

	response.SuccessWithData(c, "success", map[string]any{
		"list": list,
	})
}
