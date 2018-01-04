package redisPlayground

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"redisPlayground"
)

type MockHttpClient struct{}

func (m *MockHttpClient) Do(url *http.Request) (*http.Response, error) {
	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
	}

	return response, nil
}

func TestSendWithValidResponse(t *testing.T) {
	req := redisPlayground.CreateReqObject("https://api.bitflyer.jp")
	httpClient := &MockHttpClient{}
	err := redisPlayground.Send(httpClient, req)

	if err != nil {
		t.Errorf("Shouldn't have received an error with a valid MockHttpClient, got %s", err)
	}
}
