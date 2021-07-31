package reflect

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type TestStruct struct {
	A int    `json:"a"`
	B string `json:"b"`
	c string `json:"c"`
}

// 通过反射修改非导出字段
func TestChangeNotExportFiled(t *testing.T) {
	var r TestStruct
	// 获取字段对象
	v := reflect.ValueOf(&r).Elem().FieldByName("c")
	// 构建指向该字段的可寻址（addressable）反射对象
	rv := reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
	// 设置值
	fv := reflect.ValueOf("pibigstar")
	rv.Set(fv)

	t.Logf("%+v", r)
}

// 根据反射判断字段类型
func TestInterface(t *testing.T) {
	var value interface{}
	value = "pibigstar"
	switch value.(type) {
	case string:
		v, ok := value.(string)
		if ok {
			t.Logf("String ==> %s \n", v)
		}
	case map[string]string:
		v, ok := value.(map[string]string)
		if ok {
			t.Logf("Map ==> %v \n", v)
		}
	default:
		bs, _ := json.Marshal(value)
		t.Logf("Others ==> %s \n", string(bs))
	}
}

// 反射基本操作
func TestReflect(t *testing.T) {

	var str = "hello world"

	v := reflect.ValueOf(str)
	// 获取值
	t.Log("value:", v)
	t.Log("value:", v.String())

	// 获取类型
	t.Log("type:", v.Type())
	t.Log("kind:", v.Kind())

	// 修改值
	// 判断是否可以修改
	canSet := v.CanSet()
	t.Log("can set:", canSet)

	// 如果想修改其值，必须传递的是指针
	v = reflect.ValueOf(&str)
	v = v.Elem()
	v.SetString("new world") // 不可以直接修改

	t.Log("value:", v)

	// 通过反射修改结构体
	test := TestStruct{A: 23, B: "hello world"}
	s := reflect.ValueOf(&test).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		t.Logf("%s: Type ==>%s Value==> %v \n", typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("new world")
	t.Logf("%+v", test)
}

// 获取tag
func TestGetTag(t *testing.T) {
	s := TestStruct{}
	rt := reflect.TypeOf(s)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		fmt.Println(f.Tag.Get("json"))
	}
}
