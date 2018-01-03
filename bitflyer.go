package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type HttpClient interface {
	Get(string) (*http.Response, error)
	Do(req *http.Request) (*http.Response, error)
}

type MockHttpClient struct{}

func (m *MockHttpClient) Do(url string) (*http.Response, error) {
	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
	}

	return response, nil
}

// func TestSendWithValidResponse(t *testing.T) {
// 	httpClient := &MockHttpClient{}
// 	err := send(httpClient, "IT_JUST_WORKS!")

// 	if err != nil {
// 		t.Errorf("Shouldn't have received an error with a valid MockHttpClient, got %s", err)
// 	}
// }

var host = "https://api.bitflyer.jp"

func main() {
	req := createReqObject(host)
	client := &http.Client{}

	send(client, req)
}

func createReqObject(uri string) *http.Request {
	key := os.Getenv("BFkey")
	secret := os.Getenv("BFSecret")
	method := "GET"
	path := "/v1/me/getbalance"

	ts := strconv.FormatInt(time.Now().Unix(), 10)
	text := ts + method + path

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(text))
	sign := hex.EncodeToString(hash.Sum(nil))

	req, err := http.NewRequest(method, uri+path, nil)
	if err != nil {
		// TODO 例外処理
	}

	req.Header.Set("content-type", "application/json; charset=UTF-8")
	req.Header.Set("ACCESS-KEY", key)
	req.Header.Set("ACCESS-TIMESTAMP", ts)
	req.Header.Set("ACCESS-SIGN", sign)

	return req
}

func send(client HttpClient, req *http.Request) error {
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	return err
}

// resp, _ := client.Do(req)
// defer resp.Body.Close()

// byteArray, _ := ioutil.ReadAll(resp.Body)
// fmt.Println(string(byteArray))
