package oauth

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()

	os.Exit(m.Run())

}

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

func TestGetAccessTokenInvalidRestClientResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodGet,
		URL:          "http://localhost:8080/oauth/access_token/AbC123",
		ReqBody:      ``,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	accessToken, err := getAccessToken("AbC123")
	assert.Nil(t, accessToken)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to get access token", err.Message)
}
