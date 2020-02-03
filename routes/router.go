package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lelinu/loki/config"
	"github.com/lelinu/loki/controllers/kyb"
	"github.com/lelinu/loki/controllers/ping"
	"github.com/lelinu/loki/providers/check_provider"
	"github.com/lelinu/loki/services"
)

func BuildRouter() *gin.Engine{

	router := gin.Default()

	router.Use(
		gin.Recovery(),
	)

	rg := router.Group(config.GetApiVersionUrl())
	loadControllers(rg)

	return router
}

func loadControllers(rg *gin.RouterGroup){

	kybService := services.NewKybService(check_provider.NewProvider(nil))

	kyb.NewController(rg, kybService)
	ping.NewController(rg)
}