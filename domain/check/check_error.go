package check

import "encoding/xml"

type ErrorResponseEnvelope struct {
	XMLName xml.Name  `xml:"Envelope"`
	Body ErrorResponseBody
}

type ErrorResponseBody struct {
	XMLName     xml.Name  `xml:"Body"`
	ErrorResponse ErrorResponse
}

type ErrorResponse struct {
	XMLName   xml.Name `xml:"SoapException"`
	Text      string   `xml:",chardata"`
	Code 	  int  	   `xml:"Code"`
	Message   string   `xml:"Message"`
}