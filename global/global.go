package global

import (
	"awesomeProject/config"

	"k8s.io/client-go/kubernetes"
)

var (
	CONF          config.Server
	KubeClientSet *kubernetes.Clientset
)
