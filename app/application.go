package app

import (
	"github.com/lelinu/loki/config"
	"github.com/lelinu/loki/routes"
)

func StartApp(){
	router := routes.BuildRouter()
	if err := router.Run(config.GetApiPort()); err != nil{
		panic(err)
	}
}