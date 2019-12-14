package runtime

import (
	"runtime/debug"
	"testing"
)

func TestPrintTrace(t *testing.T) {
	test1()
}

func test1() {
	test2()
}

func test2() {
	debug.PrintStack()
}
