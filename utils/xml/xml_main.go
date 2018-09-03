package main

import (
	"encoding/xml"
	"fmt"
	"log"
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

func main() {

	addr := Address{City: "上海", State: "中国"}

	p := &Person{Id: 13, FirstName: "Pi", LastName: "bigstar", Age: 20, Height: 17.8, Married: true, Address: addr, Comment: "我是注释"}

	personXml, err := xml.Marshal(p)
	if err != nil {
		log.Println("Error when marshal", err.Error())
	}
	fmt.Println(string(personXml))
	// xml 文本转 对象
	//xml.Unmarshal()
}

/*
输出
<person id="13">
	<name>
		<first>Pi</first>
		<last>bigstar</last>
	</name>
	<age>20</age>
	<height>17.8</height>
	<Married>true</Married>
	<City>上海</City>
	<State>中国</State>
	<!--我是注释-->
</person>
*/
