package stack

import "testing"

func TestPrintStack(t *testing.T) {
	stack := GetStack()
	t.Log(stack)
}
