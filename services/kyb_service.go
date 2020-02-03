package services

import (
	"github.com/lelinu/loki/config"
	"github.com/lelinu/loki/domain/check"
	"github.com/lelinu/loki/domain/kyb"
	"github.com/lelinu/loki/providers/check_provider"
	"github.com/lelinu/loki/utils/errors"
	"strings"
)

type kybService struct {
	kybProvider *check_provider.Provider
}

type kybServiceInterface interface {
	AcknowledgeDecision(request *kyb.AcknowledgeDecisionRequest) (*kyb.AcknowledgeDecisionResponse, errors.ApiErrorInterface)
}

var (
	KybService kybServiceInterface
)

func NewKybService(kybProvider *check_provider.Provider) *kybService {
	return &kybService{kybProvider: kybProvider}
}

func (k *kybService) AcknowledgeDecision(req *kyb.AcknowledgeDecisionRequest) (*kyb.AcknowledgeDecisionResponse, errors.ApiErrorInterface) {
	req.Uuid = strings.TrimSpace(req.Uuid)
	if req.Uuid == "" {
		return nil, errors.NewBadRequestError("Invalid Uuid")
	}

	request := &check.AcknowledgeDecisionRequest{
		Username: config.GetCheckUsername(),
		Password: config.GetCheckPassword(),
		Uid:      req.Uuid,
	}

	response, err := k.kybProvider.AcknowledgeDecision(request)
	if err != nil {
		return nil, errors.NewApiError(err.Code, err.Message)
	}

	result := &kyb.AcknowledgeDecisionResponse{
		Result: response.AcknowledgeDecisionResult,
	}

	return result, nil
}
