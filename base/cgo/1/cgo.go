package main

/*
#cgo LDFLAGS: -L/usr/local/lib

#include <stdio.h>
#include <stdlib.h>
// CGO会保留C代码块中的宏定义
#define REPEAT_LIMIT 3

// 自定义结构体
typedef struct{
    int repeat_time;
    char* str;
}blob;

// 自定义函数
int SayHello(blob* pblob) {
    for ( ;pblob->repeat_time < REPEAT_LIMIT; pblob->repeat_time++){
        puts(pblob->str);
    }
    return 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	cblob := C.blob{} // 在GO程序中创建的C对象，存储在Go的内存空间
	cblob.repeat_time = 0

	cblob.str = C.CString("Hello, World\n") // C.CString 会在C的内存空间申请一个C语言字符串对象，再将Go字符串拷贝到C字符串

	// &cblob 取C语言对象cblob的地址
	ret := C.SayHello(&cblob)

	fmt.Println("ret", ret)
	fmt.Println("repeat_time", cblob.repeat_time)

	// C.CString 申请的C空间内存不会自动释放，需要显示调用C中的free释放
	C.free(unsafe.Pointer(cblob.str))
}
