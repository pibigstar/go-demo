package main

//#include <hello.h>
import "C"
import "fmt"

// 通过go实现C函数，并导出
// //export SayHello 指令将 Go 语言实现的 SayHello 函数导出为 C 语言函

//export SayHello
func SayHello(str *C.char) {
	fmt.Println("Go.....")
	fmt.Println(C.GoString(str))
}
