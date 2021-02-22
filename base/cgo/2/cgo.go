package main

/*
#include "hello.c"
int SayHello();
*/
import "C"
import (
	"fmt"
)

func main() {
	ret := C.SayHello()
	fmt.Println(ret)
}
