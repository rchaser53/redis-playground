// tryUsingXmlUnmarshal is kinda training
import (
	"encoding/xml"
	"fmt"
	"log"
)

type Head struct {
	Div, Script string
}

type Html struct {
	Head Head
}

func tryUsingXmlUnmarshal() {
	data := new(Html)
	bytes := []byte(`
		<Html>
			<Head><Script>const</Script><Div>abc</Div></Head>
  	</Html>`)

	if err := xml.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
