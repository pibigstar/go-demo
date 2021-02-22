package main

//#include "hello.h"
import "C"

func main() {
	// Go 程序先调用 C 的 SayHello 接口，由于 SayHello 接口链接在 Go 的实现上，又调到 Go
	C.SayHello(C.CString("Hello World"))
}
