package jsonpath

import (
	"encoding/json"
	"testing"

	"github.com/oliveagle/jsonpath"
)

func TestJsonPath(t *testing.T) {
	var jsonData interface{}
	if err := json.Unmarshal([]byte(ss), &jsonData); err != nil {
		t.Error(err)
	}

	// 取json里某个key的值
	res, err := jsonpath.JsonPathLookup(jsonData, "$.expensive")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)

	// 取json里某个key的值,如果是结构体会返回map
	res, err = jsonpath.JsonPathLookup(jsonData, "$.user")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)

	// 取book数组最后一组的isbn值
	res, err = jsonpath.JsonPathLookup(jsonData, "$.store.book[-1].isbn")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)

	// 取book数组下标0,1的price值
	res, err = jsonpath.JsonPathLookup(jsonData, "$.store.book[0,1].price")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)

	// 取book数组里 price 大于 10 的 price值
	res, err = jsonpath.JsonPathLookup(jsonData, "$.store.book[?(@.price > 10)].title")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)

	// 取book数组里大于该json里expensive变量的price值
	pat, _ := jsonpath.Compile(`$.store.book[?(@.price < $.expensive)].price`)
	res, err = pat.Lookup(jsonData)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

var ss = `{
    "expensive": 10,
	"user": {
		"name": "pibigstar",
		"age": 10
	},
    "store": {
        "book": [
            {
                "category": "reference",
                "author": "Nigel Rees",
                "title": "Sayings of the Century",
                "price": 8.95
            },
            {
                "category": "fiction",
                "author": "Evelyn Waugh",
                "title": "Sword of Honour",
                "price": 12.99
            },
            {
                "category": "fiction",
                "author": "Herman Melville",
                "title": "Moby Dick",
                "isbn": "0-553-21311-3",
                "price": 8.99
            },
            {
                "category": "fiction",
                "author": "J. R. R. Tolkien",
                "title": "The Lord of the Rings",
                "isbn": "0-395-19395-8",
                "price": 22.99
            }
        ],
        "bicycle": {
            "color": "red",
            "price": 19.95
        }
    }
}`
