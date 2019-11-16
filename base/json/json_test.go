package json

import (
	"testing"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	strJSON string
	user    *User
	m       map[string]interface{}
)

func init() {
	user = &User{
		ID:   "1",
		Name: "pibigstar",
	}
	m = make(map[string]interface{})
	strJSON = `{"id":"1","name":"pibigstar"}`
}

func TestStructToJson(t *testing.T) {
	strJSON = StructToJSON(user)
	t.Log(strJSON)
}

func TestJsonToStruct(t *testing.T) {
	var user User
	err := JSONToStruct(strJSON, &user)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", user)
}

func TestMapToJson(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "pibigstar"
	m["id"] = "1"
	strJSON = MapToJSON(m)
	t.Log(strJSON)
}

func TestJsonToMap(t *testing.T) {
	m := make(map[string]interface{})
	err := JSONToMap(strJSON, m)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", m)
}

func TestMapToStruct(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "pibigstar"
	m["id"] = "1"
	var user User
	err := MapToStruct(m, &user)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", user)
}

func TestStructToMap(t *testing.T) {
	user := &User{
		ID:   "1",
		Name: "pibigstar",
	}
	m := make(map[string]interface{})
	err := StructToMap(user, m)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", m)
}
