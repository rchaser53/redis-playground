package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Head struct {
	Div, Script string
}

type Html struct {
	Head Head
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
	data := new(Html)

	// bytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	bytes := []byte(`
		<Html>
			<Head><Script>const</Script><Div>abc</Div></Head>
  	</Html>`)

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
