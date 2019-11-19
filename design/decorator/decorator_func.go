package decorator

import (
	"fmt"
	"net/http"
	"reflect"
)

// 方法的装饰器模式, 包装某个方法
// http 鉴权
func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("token")
		if name == "pi" {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}

	}
}

func f(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// 用 reflection 机制写的一个比较通用的修饰器
// https://coolshell.cn/articles/17929.html
func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}
