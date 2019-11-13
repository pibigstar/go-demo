package json

import (
	"encoding/json"
	"github.com/goinggo/mapstructure"
)

func StructToJson(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	return string(bs)
}

func JsonToStruct(str string, result interface{}) error {
	return json.Unmarshal([]byte(str), result)
}

func MapToJson(params map[string]interface{}) string {
	bs, _ := json.Marshal(params)
	return string(bs)
}

func JsonToMap(str string, result map[string]interface{}) error {
	return json.Unmarshal([]byte(str), &result)
}

func MapToStruct(params map[string]interface{}, result interface{}) error {
	err := mapstructure.Decode(params, result)
	return err
}

func StructToMap(obj interface{}, result map[string]interface{}) error {
	j, err := json.Marshal(obj)
	err = json.Unmarshal(j, &result)
	return err
}
