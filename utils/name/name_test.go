package name

import (
	"testing"
)

func TestName(t *testing.T) {
	t.Log("name:", GenerateUserName(2))
	t.Log("name:", GenerateUserName(3))
}
