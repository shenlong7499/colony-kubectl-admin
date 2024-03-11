package main

import (
	"awesomeProject/global"
	"awesomeProject/initiallize"
)

func main() {
	routers := initiallize.Routers()
	initiallize.Viper()
	initiallize.K8S()
	panic(routers.Run(global.CONF.System.Addr))
}
