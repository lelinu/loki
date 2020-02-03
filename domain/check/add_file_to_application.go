package check

import "encoding/xml"

type AddFileToApplicationRequest struct {
	XMLName   xml.Name `xml:"AddFileToApplicationRequest"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Username  string   `xml:"username"`
	Password  string   `xml:"password"`
	Uid       string   `xml:"uid"`
	Filename  string   `xml:"filename"`
	Filebytes []byte   `xml:"filebytes"`
}

type AddFileToApplicationResponse struct {
	XMLName                   xml.Name `xml:"AddFileToApplicationResponse"`
	Text                      string   `xml:",chardata"`
	Xmlns                     string   `xml:"xmlns,attr"`
	AcknowledgeDecisionResult int64    `xml:"AcknowledgeDecisionResult"`
}
