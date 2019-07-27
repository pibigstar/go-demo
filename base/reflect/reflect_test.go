package reflect

import (
	"reflect"
	"testing"
)

type ReflectTest struct {
	A int
	B string
}

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
	test := ReflectTest{23, "hello world"}
	s := reflect.ValueOf(&test).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		t.Logf("%s %s = %v\n", typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("new world")
	t.Logf("%+v", test)
}
