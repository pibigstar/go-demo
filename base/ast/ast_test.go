package ast

import "testing"

func TestGenConstComments(t *testing.T) {
	genConstComment("example/code.go", "example/code_msg.go")
}
