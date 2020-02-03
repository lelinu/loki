package config

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestGetApiVersionUrl(t *testing.T){
	assert.EqualValues(t, "1.0.0", GetApiVersionUrl())
}

func TestGetApiPort(t *testing.T){
	assert.EqualValues(t, "8080", GetApiPort())
}

func TestGetCheckUsername(t *testing.T) {
	assert.EqualValues(t, "", GetCheckUsername())
}

func TestGetCheckPassword(t *testing.T) {
	assert.EqualValues(t, "", GetCheckUsername())
}

func TestGetCheckUrlBase(t *testing.T) {
	checkUrlBase = url.URL{
		Scheme: "https",
		Host:   "localhost:44300",
		Path:   "Service.asmx",
	}
	assert.EqualValues(t, checkUrlBase, GetCheckUrlBase())
}