package js

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"testing"
)

// 使用Go来解释执行Js代码

func TestOttO(t *testing.T) {
	vm := otto.New()

	// 给某个变量赋值
	err := vm.Set("b", "test")
	if err != nil {
		t.Error(err)
	}

	vm.Run(`
    a = 2 + 2;
    console.log("a: " + a + ", b: " + b);
`)
	// 获取某个变量的值
	if v, err := vm.Get("a"); err == nil {
		if vInt, err := v.ToInteger(); err == nil {
			t.Log(vInt)
		}
	}

	// 使用js中的表达式
	vm.Set("c", "pibigstar")
	if value, err := vm.Run(`c.length`); err == nil {
		t.Log(value)
	}
}

// 调用JS文件中函数
func TestOttOFromJs(t *testing.T) {
	vm := otto.New()
	script, err := vm.Compile("test.js", nil)
	if err != nil {
		t.Error(err)
	}
	vm.Run(script)

	// 调用某个函数
	if v, err := vm.Call("test", nil, "pibigstar"); err == nil {
		t.Log(v)
	}
}

// 用Go来完成Js中某个函数
func TestOttOWithGo(t *testing.T) {
	vm := otto.New()
	script, err := vm.Compile("test.js", nil)
	if err != nil {
		t.Error(err)
	}
	// 用Go来完成Js中某个函数
	vm.Set("sayHello", func(call otto.FunctionCall) otto.Value {
		s := fmt.Sprintf("Hello, %s", call.Argument(0).String())
		value, _ := otto.ToValue(s)
		return value
	})
	if result, err := vm.Run(script); err == nil {
		t.Log(result)
	}
}
