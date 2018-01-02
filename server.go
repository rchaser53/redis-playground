package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/antonholmquist/jason"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type CL struct {
	Quotes Quotes
	Source string
}

type Quotes struct {
	USDJPY float64
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
func tryToUseGetJson() {
	url := "http://apilayer.net/api/live?access_key=" + os.Getenv("CLkey")
	var cl CL
	err := getJson(url, &cl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%f", cl.Quotes.USDJPY)
}

func main() {
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
	println(string(j))
}

// if b, err := ioutil.ReadAll(resp.Body); err == nil {
// 	return string(b)
// }
