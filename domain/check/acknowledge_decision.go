package check

import "encoding/xml"

type AcknowledgeDecisionRequest struct {
	XMLName  xml.Name `xml:"AcknowledgeDecision"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
	Uid      string   `xml:"uid"`
}

type AcknowledgeDecisionEnvelope struct {
	XMLName xml.Name  `xml:"Envelope"`
	Body AcknowledgeDecisionBody
}

type AcknowledgeDecisionBody struct {
	XMLName     xml.Name  `xml:"Body"`
	AcknowledgeDecisionResponse AcknowledgeDecisionResponse
}

type AcknowledgeDecisionResponse struct {
	XMLName                   xml.Name `xml:"AcknowledgeDecisionResponse"`
	Text                      string   `xml:",chardata"`
	Xmlns                     string   `xml:"xmlns,attr"`
	AcknowledgeDecisionResult uint64    `xml:"AcknowledgeDecisionResult"`
}

