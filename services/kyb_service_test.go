package services

import (
	"github.com/lelinu/loki/clients/soap"
	"github.com/lelinu/loki/config"
	"github.com/lelinu/loki/domain/kyb"
	"github.com/lelinu/loki/providers/check_provider"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	soapClient *soap.Client
)

func TestMain(m *testing.M) {
	soapClient = soap.NewClient(nil, config.GetCheckUrlBase())
	soapClient.SetMock(true)

	os.Exit(m.Run())
}

func TestAcknowledgeDecisionInvalidRequest(t *testing.T){
	provider := check_provider.NewProvider(soapClient)
	kybService := NewKybService(provider)

	req := kyb.AcknowledgeDecisionRequest{}
	resp, err := kybService.AcknowledgeDecision(&req)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, "Invalid Uuid", err.Message())
}

func TestAcknowledgeDecisionCheckError(t *testing.T){

}

func TestAcknowledgeDecisionCheckSuccess(t *testing.T){

}