package json

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"reflect"
	"strings"
)

func StructToJSON(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	return string(bs)
}

func JSONToStruct(str string, result interface{}) error {
	return json.Unmarshal([]byte(str), result)
}

func MapToJSON(params map[string]interface{}) string {
	bs, _ := json.Marshal(params)
	return string(bs)
}

func JSONToMap(str string, result map[string]interface{}) error {
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

// 将interface，从float64更换成int64
func MarshInterface(jsonStr string) error {
	decoder := json.NewDecoder(strings.NewReader(jsonStr))
	// 让interface{}反解析成int，而不是float64
	decoder.UseNumber()

	var user map[string]interface{}
	if err := decoder.Decode(&user); err != nil {
		return err
	}

	num := user["num"]
	fmt.Printf("%+v \n", reflect.TypeOf(num).PkgPath()+"."+reflect.TypeOf(num).Name())

	v, err := num.(json.Number).Int64()
	if err != nil {
		return err
	}
	fmt.Printf("%+v \n", v)
	return nil
}
