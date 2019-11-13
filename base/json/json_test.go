package json

import (
	"testing"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	strJson string
	user    *User
	m       map[string]interface{}
)

func init() {
	user = &User{
		ID:   "1",
		Name: "pibigstar",
	}
	m = make(map[string]interface{})
	strJson = `{"id":"1","name":"pibigstar"}`
}

func TestStructToJson(t *testing.T) {
	strJson = StructToJson(user)
	t.Log(strJson)
}

func TestJsonToStruct(t *testing.T) {
	var user User
	err := JsonToStruct(strJson, &user)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", user)
}

func TestMapToJson(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "pibigstar"
	m["id"] = "1"
	strJson = MapToJson(m)
	t.Log(strJson)
}

func TestJsonToMap(t *testing.T) {
	m := make(map[string]interface{})
	err := JsonToMap(strJson, m)
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
