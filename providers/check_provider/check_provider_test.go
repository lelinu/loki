package check_provider

import (
	"errors"
	"github.com/lelinu/loki/clients/soap"
	"github.com/lelinu/loki/config"
	"github.com/lelinu/loki/domain/check"
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

func TestVariables(t *testing.T) {
	assert.NotNil(t, actionAcknowledgeDecision)
	assert.EqualValues(t, "AcknowledgeDecision", actionAcknowledgeDecision)
}

func TestNewProviderWithSoapClient(t *testing.T) {
	provider := NewProvider(soapClient)
	assert.NotNil(t, provider)
	assert.NotNil(t, provider.soapClient)
}

func TestNewProviderWithoutSoapClient(t *testing.T) {
	provider := NewProvider(nil)
	assert.NotNil(t, provider)
	assert.NotNil(t, provider.soapClient)
}

func TestAcknowledgeDecisionErrorSoapClient(t *testing.T) {

	soapClient.ResetMocks()

	mock := soap.NewMock("https://localhost:44300/Service.asmx",
		"AcknowledgeDecision",
		http.MethodPost, nil,
		errors.New("invalid soap client response"))

	soapClient.AddMock(mock)
	provider := NewProvider(soapClient)

	response, err := provider.AcknowledgeDecision(&check.AcknowledgeDecisionRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "invalid soap client response", err.Message)
}

func TestAcknowledgeDecisionInvalidResponseBody(t *testing.T) {

	invalidCloser, _ := os.Open("lelinu")
	soapClient.ResetMocks()

	mock := soap.NewMock("https://localhost:44300/Service.asmx",
		"AcknowledgeDecision",
		http.MethodPost,
		&http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       invalidCloser,
		}, nil)

	soapClient.AddMock(mock)
	provider := NewProvider(soapClient)

	response, err := provider.AcknowledgeDecision(&check.AcknowledgeDecisionRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "Invalid Response Body", err.Message)
}

func TestAcknowledgeDecisionInvalidErrorInterface(t *testing.T) {

	soapClient.ResetMocks()

	mock := soap.NewMock("https://localhost:44300/Service.asmx",
		"AcknowledgeDecision",
		http.MethodPost,
		&http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       ioutil.NopCloser(strings.NewReader(`"message": 1`)),
		},
		nil)

	soapClient.AddMock(mock)
	provider := NewProvider(soapClient)

	response, err := provider.AcknowledgeDecision(&check.AcknowledgeDecisionRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "Invalid XML Body", err.Message)
}

func TestAcknowledgeDecisionInvalidSuccessResponse(t *testing.T) {
	soapClient.ResetMocks()

	mock := soap.NewMock("https://localhost:44300/Service.asmx",
		"AcknowledgeDecision",
		http.MethodPost,
		&http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`"message": 1`)),
		},
		nil)

	soapClient.AddMock(mock)
	provider := NewProvider(soapClient)

	response, err := provider.AcknowledgeDecision(&check.AcknowledgeDecisionRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "Invalid XML Body", err.Message)
}

func TestAcknowledgeDecisionUnauthorized(t *testing.T) {
	soapClient.ResetMocks()

	mock := soap.NewMock("https://localhost:44300/Service.asmx",
		"AcknowledgeDecision",
		http.MethodPost,
		&http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><SoapException xmlns="http://tempuri.org/"><Code>401</Code><Message>Unauthorized</Message></SoapException></soap:Body></soap:Envelope>`)),
		},
		nil)

	soapClient.AddMock(mock)
	provider := NewProvider(soapClient)

	response, err := provider.AcknowledgeDecision(&check.AcknowledgeDecisionRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Code)
}

func TestAcknowledgeDecisionValidResponse(t *testing.T) {
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
	provider := NewProvider(soapClient)

	response, err := provider.AcknowledgeDecision(&check.AcknowledgeDecisionRequest{})

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, response.AcknowledgeDecisionResult)
}
