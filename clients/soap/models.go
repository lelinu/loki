package soap

import (
	"encoding/xml"
	"net/http"
)

type mock struct{
	url string
	action string
	httpMethod string
	response *http.Response
	err error
}

func NewMock(url string, action string, httpMethod string, response *http.Response, err error) *mock{
	return &mock{
		url: url,
		action: action,
		httpMethod: httpMethod,
		response: response,
		err: err,
	}
}

type soapRequest struct {
	XMLName   xml.Name `xml:"soap:Envelope"`
	XMLNsSoap string   `xml:"xmlns:soap,attr"`
	XMLNsXSI  string   `xml:"xmlns:xsi,attr"`
	XMLNsXSD  string   `xml:"xmlns:xsd,attr"`
	Body      soapBody
}

type soapBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Payload interface{}
}

func newSoapRequest(body interface{}) *soapRequest{
	return &soapRequest{
		XMLNsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XMLNsXSD:  "http://www.w3.org/2001/XMLSchema",
		XMLNsXSI:  "http://www.w3.org/2001/XMLSchema-instance",
		Body: soapBody{
			Payload: body,
		},
	}
}