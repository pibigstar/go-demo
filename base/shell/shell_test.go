package shell

import (
	"os/exec"
	"testing"
)

func TestExecCmd(t *testing.T) {
	out, err := exec.Command("cmd", "/c", "echo", "Hello World").Output()
	if err != nil {
		t.Fatal(t)
	}
	t.Log(string(out))
}
