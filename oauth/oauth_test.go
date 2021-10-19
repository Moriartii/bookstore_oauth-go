package oauth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOauthConstant(t *testing.T) {
	if headerXPublic != "X-Public" {
		t.Error("headerXPublic should be -> X-Public")
	}
	if headerXClientId != "X-Client-Id" {
		t.Error("headerXClientId should be -> X-Client-Id")
	}
	if headerXCallerId != "X-Caller-Id" {
		t.Error("headerXCallerId should be -> X-Caller-Id")
	}
	if paramAccessTolen != "access_token" {
		t.Error("paramAccessTolen should be -> access_token")
	}
}

func TestIsPublicNilRequest(t *testing.T) {
	if IsPublic(nil) != true {
		t.Error("When request is nil ther bool return must be true")
	}
}

func TestIsPublicNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	if IsPublic(&request) != false {
		t.Error("If request without X-Public(header) it must return false")
	}

	request.Header.Add("X-Public", "true")

	if IsPublic(&request) != true {
		t.Error("If request with X-Public(header) must return true")
	}

}

func TestGetCallerIDNilRequest(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	assert.Zero(t, GetCallerID(&request))
}

func TestGetCallerIDInvalidCallerFormat(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	request.Header.Add("X-Caller-Id", "badID_STRING")
	assert.Zero(t, GetCallerID(&request))

}

func TestGetCallerIDNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	request.Header.Add("X-Caller-Id", "123")
	assert.EqualValues(t, 123, GetCallerID(&request))
}
