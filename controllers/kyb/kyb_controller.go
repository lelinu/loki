package kyb

import (
	"github.com/gin-gonic/gin"
	"github.com/lelinu/loki/domain/kyb"
	"github.com/lelinu/loki/services"
	"github.com/lelinu/loki/utils/errors"
	"net/http"
)

type (
	Controller struct {
		kybService services.KybServiceInterface
	}
)

func NewController(rg *gin.RouterGroup, kybService services.KybServiceInterface){
	controller := &Controller{kybService: kybService}

	group :=  rg.Group("kyb")
	{
		group.POST("/acknowledgedecision", controller.AcknowledgeDecision())
	}
}

func (k *Controller) AcknowledgeDecision() func(c *gin.Context) {

	return func(c *gin.Context) {
		var req kyb.AcknowledgeDecisionRequest
		if err := c.ShouldBindJSON(req); err != nil {
			apiErr := errors.NewBadRequestError("Invalid JSON request")
			c.JSON(apiErr.Status(), apiErr)
			return
		}

		result, err := k.kybService.AcknowledgeDecision(&req)
		if err != nil{
			c.JSON(err.Status(), err)
			return
		}

		c.JSON(http.StatusOK, result)
		return
	}
}
