package api

import (
	"awesomeProject/api/example"
	"awesomeProject/api/kubectl"
)

type ApiGroup struct {
	ExampleApiGroup example.ExampleApiGroup
	KubectlApiGroup kubectl.KubectlApiGroup
}

var ApiGroupApp = new(ApiGroup)
