package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Name struct {
	First, Last string
}

type Person struct {
	Name   Name
	Gender string
	Age    int
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	// resp, err := http.Get("https://github.com/rchaser53/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	data := new(Person)

	// bytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	bytes := []byte(`
		<Person>
    <Name><First>John</First><Last>Doe</Last></Name>
    <Gender>Male</Gender>
    <Age>20</Age>
  </Person>`)

	if err := xml.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	// println(data.html)

	// defer resp.Body.Close()
}

// println(string(b))

// if b, err := ioutil.ReadAll(resp.Body); err == nil {
// 	return string(b)
// }
