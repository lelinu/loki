package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Controller struct {}
)

func NewController(rg *gin.RouterGroup){
	controller := &Controller{}

	group :=  rg.Group("ping")
	{
		group.POST("/me", controller.Me())
	}
}

func (k *Controller) Me() func(c *gin.Context) {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "Hello from LOKI Api")
		return
	}
}
