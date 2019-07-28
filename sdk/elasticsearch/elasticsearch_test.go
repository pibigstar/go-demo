package elastic

import (
	"encoding/json"
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
		t.Error(err)
	}
	t.Logf("%+v", response)

	bytes, err := client.GetById("test", "user", response.Id)
	if err != nil {
		t.Error(err)
	}
	result := new(User)
	json.Unmarshal(bytes, result)
	t.Logf("%+v", result)
}
