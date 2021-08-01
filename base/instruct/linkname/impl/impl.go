package impl

import (
	"fmt"
	_ "unsafe" // for go:linkname
)

// 通过访问 outer.Linkname 就是访问该函数
//go:linkname hello go-demo/base/instruct/linkname/outer.Linkname
func hello() {
	fmt.Println("hello world")
}
