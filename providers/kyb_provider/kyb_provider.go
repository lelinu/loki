package kyb_provider

import (
	"github.com/lelinu/loki/clients/soap"
	"github.com/lelinu/loki/domain/kyb"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	urlBase = url.URL{
		Scheme: "https",
		Host:   "localhost:44300",
		Path:   "Service.asmx",
	}

	actionAcknowledgeDecision = "AcknowledgeDecision"
)

type Provider struct {
	soapClient *soap.Client
}

func NewProvider(soapClient *soap.Client) *Provider {

	if soapClient == nil {
		soapClient = soap.NewClient(nil, urlBase)
	}

	provider := &Provider{
		soapClient: soapClient,
	}

	return provider
}

func (p *Provider) AcknowledgeDecision(req *kyb.AcknowledgeDecisionRequest) (*kyb.AcknowledgeDecisionResponse, *kyb.ErrorResponse) {

	//post using soap client
	response, err := p.soapClient.Post(actionAcknowledgeDecision, req)
	if err != nil {
		return nil, &kyb.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer response.Body.Close()

	//read response body
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &kyb.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Invalid Response Body",
		}
	}

	if response.StatusCode > 299 {

		var errResponse kyb.ErrorResponseEnvelope
		if err := xml.Unmarshal(respBytes, &errResponse); err != nil {
			return nil, &kyb.ErrorResponse{
				Code:    response.StatusCode,
				Message: "Invalid XML Body",
			}
		}
		errResponse.Body.ErrorResponse.Code = response.StatusCode
		return nil, &errResponse.Body.ErrorResponse
	}

	//build acknowledge decision response
	var result kyb.AcknowledgeDecisionEnvelope
	if err := xml.Unmarshal(respBytes, &result); err != nil {

		return nil, &kyb.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Invalid XML Body",
		}
	}
	return &result.Body.AcknowledgeDecisionResponse, nil
}
