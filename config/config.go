package config

import (
	"net/url"
	"os"
)

const(
	secretCheckUsername = "SECRET_CHECK_USERNAME"
	secretCheckPassword  = "SECRET_CHECK_PASSWORD"
)

var(
	checkUsername = os.Getenv(secretCheckUsername)
	checkPassword = os.Getenv(secretCheckPassword)
	checkUrlBase = url.URL{
		Scheme: "https",
		Host:   "localhost:44300",
		Path:   "Service.asmx",
	}
)

func GetCheckUsername() string{
	return checkUsername
}

func GetCheckPassword() string{
	return checkPassword
}

func GetCheckUrlBase() url.URL{
	return checkUrlBase
}
