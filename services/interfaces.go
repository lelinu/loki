package services

import (
	"github.com/lelinu/loki/domain/kyb"
	"github.com/lelinu/loki/utils/errors"
)

type KybServiceInterface interface {
	AcknowledgeDecision(request *kyb.AcknowledgeDecisionRequest) (*kyb.AcknowledgeDecisionResponse, errors.ApiErrorInterface)
}