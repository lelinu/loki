package kyb

import "encoding/xml"

type QApp struct {
	XMLName        xml.Name             `xml:"a"`
	Text           string               `xml:",chardata"`
	Uid            string               `xml:"Uid"`
	Company        QCompany             `xml:"Company"`
	Comments       string               `xml:"Comments"`
	Status         string               `xml:"Status"`
	SubStatus      string               `xml:"Substatus"`
	ExtSubStatus   []QExtendedSubStatus `xml:"ExtSubstatus"`
	Pepcheckstatus int64                `xml:"Pepcheckstatus"`
}

type QCompany struct {
	XMLName   xml.Name    `xml:"Company"`
	Text      string      `xml:",chardata"`
	Fields    []QField    `xml:"Fields"`
	Directors []QDirector `xml:"Directors"`
	Ubos      []QUbo      `xml:"Ubos"`
	Documents []QDoc      `xml:"Documents"`
}

type QField struct {
	XMLName xml.Name `xml:"QField"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name"`
	Value   string   `xml:"value"`
}

type QDirector struct {
	XMLName   xml.Name `xml:"QDirector"`
	Text      string   `xml:",chardata"`
	Fields    []QField `xml:"Fields"`
	Documents []QDoc   `xml:"Documents"`
}

type QUbo struct {
	XMLName   xml.Name `xml:"QUbo"`
	Text      string   `xml:",chardata"`
	Fields    []QField `xml:"Fields"`
	Company   QCompany `xml:"Company"`
	Documents []QDoc   `xml:"Documents"`
}

type QDoc struct {
	XMLName   xml.Name `xml:"QDoc"`
	Text      string   `xml:",chardata"`
	Filename  string   `xml:"Filename"`
	FileBytes []byte   `xml:"FileBytes"`
	ClassType string   `xml:"ClassType"`
	DocType   string   `xml:"DocType"`
	Decisions []string `xml:"Decisions"`
}

type QExtendedSubStatus struct {
	XMLName xml.Name `xml:"ExtendedSubStatus"`
	Text    string   `xml:",chardata"`
	Value   string   `xml:"Value"`
}

type SubmitApplicationRequest struct {
	XMLName  xml.Name `xml:"SubmitApplication"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
	Uid      string   `xml:"uid"`
	QApp     QApp
}

type SubmitApplicationResponse struct {
	XMLName                 xml.Name `xml:"SubmitApplicationResponse"`
	Text                    string   `xml:",chardata"`
	Xmlns                   string   `xml:"xmlns,attr"`
	SubmitApplicationResult int64    `xml:"SubmitApplicationResult"`
}
