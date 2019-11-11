package xml

import (
	"encoding/xml"
	"log"
	"testing"
)

/**
对象转xml
*/
type Address struct {
	City  string
	State string
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}

func TestXml(t *testing.T) {
	addr := Address{City: "上海", State: "中国"}
	p := &Person{Id: 13,
		FirstName: "Pi",
		LastName:  "bigstar",
		Age:       20, Height: 17.8,
		Married: true,
		Address: addr,
		Comment: "我是注释",
	}
	// 对象转xml
	personXML, err := xml.Marshal(p)
	if err != nil {
		log.Println("Error when marshal", err.Error())
	}
	t.Log(string(personXML))
	//xml 文本转 对象
	obj := &Person{}
	xml.Unmarshal(personXML, obj)
	t.Logf("%+v", obj)
}
