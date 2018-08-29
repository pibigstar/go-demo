package json

import (
	"encoding/json"
	"fmt"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

// 对象转Json字符串 json.Marshal
func ObjectToJson() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

}

// Json字符串转对象 json.Unmarshal
func JsonToObjcet() {
	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals) // 添加一个 + 号可以输出key值
}
