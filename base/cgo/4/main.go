package main

/*
#include "hello.h"
*/
import "C"
import (
	"fmt"
)

func main() {
	ret := C.SayHello()
	fmt.Println(ret)
}
