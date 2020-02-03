package services

import (
	"errors"
	"github.com/lelinu/loki/clients/soap"
	"github.com/lelinu/loki/config"
	"github.com/lelinu/loki/domain/kyb"
	"github.com/lelinu/loki/providers/check_provider"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	soapClient.ResetMocks()

	mock := soap.NewMock("https://localhost:44300/Service.asmx",
		"AcknowledgeDecision",
		http.MethodPost, nil,
		errors.New("an error had occurred in check provider"))

	soapClient.AddMock(mock)

	provider := check_provider.NewProvider(soapClient)
	kybService := NewKybService(provider)

	req := kyb.AcknowledgeDecisionRequest{
		Uuid: "3df62934-843a-4216-8970-7a583f8092ee",
	}

	resp, err := kybService.AcknowledgeDecision(&req)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "an error had occurred in check provider", err.Message())
}

func TestAcknowledgeDecisionCheckSuccess(t *testing.T){
	soapClient.ResetMocks()

	mock := soap.NewMock("https://localhost:44300/Service.asmx",
		"AcknowledgeDecision",
		http.MethodPost,
		&http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><AcknowledgeDecisionResponse xmlns="http://tempuri.org/"><AcknowledgeDecisionResult>1</AcknowledgeDecisionResult></AcknowledgeDecisionResponse></soap:Body></soap:Envelope>`)),
		},
		nil)

	soapClient.AddMock(mock)

	provider := check_provider.NewProvider(soapClient)
	kybService := NewKybService(provider)

	req := kyb.AcknowledgeDecisionRequest{
		Uuid: "3df62934-843a-4216-8970-7a583f8092ee",
	}

	resp, err := kybService.AcknowledgeDecision(&req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, resp.Result)
}