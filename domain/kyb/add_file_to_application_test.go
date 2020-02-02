package kyb

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddFileToApplicationRequest(t *testing.T){

	//setup
	req := &AddFileToApplicationRequest{
		XMLName:  xml.Name{},
		Xmlns:    "http://tempuri.org/",
		Username: "12345",
		Password: "password",
		Uid:      "54321",
		Filename: "filename",
		Filebytes: make([]byte, 1),
	}

	//execute
	bytes, err := xml.Marshal(req)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	//unmarshal result
	var target = &AddFileToApplicationRequest{}
	err = xml.Unmarshal(bytes, target)

	//Assert that error is still null
	assert.Nil(t, err)
	assert.NotNil(t, target)
	assert.EqualValues(t, "12345", target.Username)
	assert.EqualValues(t, "password", target.Password)
	assert.EqualValues(t, "54321", target.Uid)
	assert.EqualValues(t, "filename", target.Filename)
	assert.NotNil(t, target.Filebytes)
}