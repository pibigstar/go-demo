package shell

import (
	"os/exec"
	"testing"
)

func TestExecCmd(t *testing.T) {
	// Windows
	out, err := exec.Command("cmd", "/c", "echo", "Hello World").Output()
	if err != nil {
		t.Fatal(t)
	}
	//Linux
	out, err = exec.Command("echo", "Hello World").Output()
	if err != nil {
		t.Fatal(t)
	}
	t.Log(string(out))
}
