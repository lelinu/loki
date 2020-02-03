package check

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAcknowledgeDecisionRequest(t *testing.T){

	//setup
	req := &AcknowledgeDecisionRequest{
		XMLName:  xml.Name{},
		Xmlns:    "http://tempuri.org/",
		Username: "12345",
		Password: "password",
		Uid:      "54321",
	}

	//execute
	bytes, err := xml.Marshal(req)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	//unmarshal result
	var target = &AcknowledgeDecisionRequest{}
	err = xml.Unmarshal(bytes, target)

	//Assert that error is still null
	assert.Nil(t, err)
	assert.NotNil(t, target)
	assert.EqualValues(t, "12345", target.Username)
	assert.EqualValues(t, "password", target.Password)
	assert.EqualValues(t, "54321", target.Uid)
}

func TestAcknowledgeDecisionResponse(t *testing.T){
	//setup
	resp := &AcknowledgeDecisionResponse{
		XMLName:                   xml.Name{},
		Xmlns:                     "http://tempuri.org/",
		AcknowledgeDecisionResult: 0,
	}

	//execute
	bytes, err := xml.Marshal(resp)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	//unmarshal result
	var target = &AcknowledgeDecisionResponse{}
	err = xml.Unmarshal(bytes, target)

	//Assert that error is still null
	assert.Nil(t, err)
	assert.NotNil(t, target)
	assert.EqualValues(t, 0, target.AcknowledgeDecisionResult)
}
