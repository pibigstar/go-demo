package shell

import (
	"os/exec"
	"runtime"
	"testing"
)

func TestExecCmd(t *testing.T) {
	sysType := runtime.GOOS
	if sysType == "linux" {
		//Linux
		out, err := exec.Command("echo", "Hello World").Output()
		if err != nil {
			t.Error(err)
		}
		t.Log(string(out))
	}

	if sysType == "windows" {
		// Windows
		out, err := exec.Command("cmd", "/c", "echo", "Hello World").Output()
		if err != nil {
			t.Error(err)
		}
		t.Log(string(out))
	}
}
