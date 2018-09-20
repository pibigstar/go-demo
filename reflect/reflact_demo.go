package main

import (
  "reflect"
  "fmt"
)

type T struct {
  A int
  B string
}

func main() {

  var str = "hello world"

  v := reflect.ValueOf(str)
  // 获取值
  fmt.Println("value:", v)
  fmt.Println("value:", v.String())

  // 获取类型
  fmt.Println("type:", v.Type())
  fmt.Println("kind:", v.Kind())

  // 修改值
  // 判断是否可以修改
  canSet := v.CanSet()
  fmt.Println("can set:",canSet)

  // 如果想修改其值，必须传递的是指针
  v = reflect.ValueOf(&str)
  v = v.Elem()
  v.SetString("new world") // 不可以直接修改

  fmt.Println("value:", v)


  // 通过反射修改结构体
  t := T{23, "hello world"}
  s := reflect.ValueOf(&t).Elem()
  typeOfT := s.Type()
  for i := 0; i < s.NumField(); i++ {
    f := s.Field(i)
    fmt.Printf("%s %s = %v\n", typeOfT.Field(i).Name, f.Type(), f.Interface())
  }
  s.Field(0).SetInt(77)
  s.Field(1).SetString("new world")
  fmt.Printf("%+v", t)

}





