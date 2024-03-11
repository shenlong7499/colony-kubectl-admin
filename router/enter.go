package router

import (
	"awesomeProject/router/example"
	"awesomeProject/router/kubectl"
)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouterGroup
	KubectlRouterGroup kubectl.KubectlRouterGroup
}

var RouterGroupApp = new(RouterGroup)
