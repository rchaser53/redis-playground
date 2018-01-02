package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type CL struct {
	Quotes Quotes ``
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

func main() {
	url := "http://apilayer.net/api/live?access_key=" + os.Getenv("CLkey")
	var cl CL
	err := getJson(url, &cl)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%f\n", cl.quotes.USDJPY)
	fmt.Printf("%f", cl.Quotes.USDJPY)

	// defer resp.Body.Close()
}

// println(string(b))

// if b, err := ioutil.ReadAll(resp.Body); err == nil {
// 	return string(b)
// }
