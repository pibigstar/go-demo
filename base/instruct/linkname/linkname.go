package main

import (
	"go-demo/base/instruct/linkname/outer"
)

// 原本 impl 里面的 hello 函数是非导出的
// 通过 //go:linkname 变相的调用了 impl 包里非导出函数

func main() {
	outer.Linkname()
}
