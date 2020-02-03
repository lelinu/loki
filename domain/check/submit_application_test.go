package check

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubmitApplicationRequest(t *testing.T){

	xmlns := "http://tempuri.org/"

	//setup
	req := &SubmitApplicationRequest{
		Xmlns:    xmlns,
		Username: "12345",
		Password: "password",
		Uid:      "54321",
		QApp:     QApp{
			Uid:            "11111",
			Company:        QCompany{
				Fields:    []QField{
					{
						Name:    "CompName",
						Value:   "KnowMeNow",
					},
				},
				Directors: []QDirector{
					{
						Fields:    []QField{
							{
								Name:    "Name",
								Value:   "Director 1",
							},
						},
						Documents: nil,
					},
				},
				Ubos:      nil,
				Documents: nil,
			},
			Comments:       "comment",
			Status:         "status",
			SubStatus:      "sub status",
			Pepcheckstatus: 0,
		},
	}

	//execute
	bytes, err := xml.Marshal(req)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	//unmarshal result
	var target = &SubmitApplicationRequest{}
	err = xml.Unmarshal(bytes, target)

	//Assert that error is still null
	assert.Nil(t, err)
	assert.NotNil(t, target)
	assert.EqualValues(t, "12345", target.Username)
	assert.EqualValues(t, "password", target.Password)
	assert.EqualValues(t, "54321", target.Uid)
	assert.EqualValues(t, "comment", target.QApp.Comments)
	assert.EqualValues(t, "status", target.QApp.Status)
	assert.EqualValues(t, "sub status", target.QApp.SubStatus)
	assert.EqualValues(t, 0, target.QApp.Pepcheckstatus)
}

func TestSubmitApplicationResponse(t *testing.T){
	//setup
	resp := &SubmitApplicationResponse{
		XMLName:                   xml.Name{},
		Xmlns:                     "http://tempuri.org/",
		SubmitApplicationResult: 0,
	}

	//execute
	bytes, err := xml.Marshal(resp)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	//unmarshal result
	var target = &SubmitApplicationResponse{}
	err = xml.Unmarshal(bytes, target)

	//Assert that error is still null
	assert.Nil(t, err)
	assert.NotNil(t, target)
	assert.EqualValues(t, 0, target.SubmitApplicationResult)
}

