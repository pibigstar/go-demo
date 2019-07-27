package elastic

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestElasticSearch(t *testing.T) {

	user := &User{Name: "pibigstar", Age: 18}

	response, err := client.Insert("test", "user", user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v", response)

	bytes, err := client.GetById("test", "user", response.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
	result := new(User)
	json.Unmarshal(bytes, result)
	fmt.Printf("%+v", result)
}
