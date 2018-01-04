package redisPlayground

import (
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
	Do(req *http.Request) (*http.Response, error)
}

var host = "https://api.bitflyer.jp"

func CreateReqObject(uri string) *http.Request {
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

func Send(client HttpClient, req *http.Request) error {
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	return err
}

// func main() {
// 	req := createReqObject(host)
// 	client := &http.Client{}

// 	send(client, req)
// }

// resp, _ := client.Do(req)
// defer resp.Body.Close()

// byteArray, _ := ioutil.ReadAll(resp.Body)
// fmt.Println(string(byteArray))
