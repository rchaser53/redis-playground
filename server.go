package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/antonholmquist/jason"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type CL struct {
	Quotes Quotes `json:"quotes"`
	Source string `json:"source"`
}

type Quotes struct {
	USDJPY float64 `json:"usdJpy"`
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// tryToUseGetJson is suburi
func tryToUseGetJson() Quotes {
	url := "http://apilayer.net/api/live?access_key=" + os.Getenv("CLkey")
	var cl CL
	err := getJson(url, &cl)

	if err != nil {
		log.Fatal(err)
	}

	return cl.Quotes
}

// tryToUseJason is suburi
func tryToUseJason() string {
	url := "http://apilayer.net/api/live?access_key=" + os.Getenv("CLkey")

	resp, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	v, err := jason.NewObjectFromBytes(b)

	a, err := v.GetObject("quotes")
	j, err := json.Marshal(a)
	return string(j)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// respStr := tryToUseJason()
	respBuffer, err := json.Marshal(tryToUseGetJson())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(respBuffer), html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":5000", nil)
}
