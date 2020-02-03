package check_provider

import (
	"github.com/lelinu/loki/clients/soap"
	"github.com/lelinu/loki/config"
	"github.com/lelinu/loki/domain/check"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

var (
	actionAcknowledgeDecision = "AcknowledgeDecision"
)

type Provider struct {
	soapClient *soap.Client
}

func NewProvider(soapClient *soap.Client) *Provider {

	if soapClient == nil {
		soapClient = soap.NewClient(nil, config.GetCheckUrlBase())
	}

	provider := &Provider{
		soapClient: soapClient,
	}

	return provider
}

func (p *Provider) AcknowledgeDecision(req *check.AcknowledgeDecisionRequest) (*check.AcknowledgeDecisionResponse, *check.ErrorResponse) {

	//post using soap client
	response, err := p.soapClient.Post(actionAcknowledgeDecision, req)
	if err != nil {
		return nil, &check.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer response.Body.Close()

	//read response body
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &check.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Invalid Response Body",
		}
	}

	if response.StatusCode > 299 {

		var errResponse check.ErrorResponseEnvelope
		if err := xml.Unmarshal(respBytes, &errResponse); err != nil {
			return nil, &check.ErrorResponse{
				Code:    response.StatusCode,
				Message: "Invalid XML Body",
			}
		}
		errResponse.Body.ErrorResponse.Code = response.StatusCode
		return nil, &errResponse.Body.ErrorResponse
	}

	//build acknowledge decision response
	var result check.AcknowledgeDecisionEnvelope
	if err := xml.Unmarshal(respBytes, &result); err != nil {

		return nil, &check.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Invalid XML Body",
		}
	}
	return &result.Body.AcknowledgeDecisionResponse, nil
}
