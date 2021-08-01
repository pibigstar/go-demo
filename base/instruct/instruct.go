package instruct

// 禁止内联
//go:noinline
func Test1() {

}

// 禁止进行竞态检测
//go:norace
func Test2() {

}

// 禁止编译器对其做逃逸分析
//go:noescape
func Test3() {

}

// 禁止内联
//go:noinline
func Test4() {

}
