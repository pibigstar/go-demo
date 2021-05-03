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

func TestExecKillChildProc(t *testing.T) {
	cmd := exec.Command("curl", "-o", "test.tar.gz", "http://test.tar.gz")
	// 设置当主进程退出时，将子进程也退出，仅linux系统支持
	//cmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGKILL}
	err := cmd.Start()
	if err != nil {
		t.Error(err)
	}
}

// 测试pipe，一个命令的输出是另一个命令的输入
func TestPipe(t *testing.T) {
	cmdCat := exec.Command("cat", "main.go")
	cmdWC := exec.Command("wc", "-l")
	data, err := pipeCommands(cmdCat, cmdWC)
	if err != nil {
		t.Error(err)
	}
	t.Logf("output: %s", data)
}

func pipeCommands(commands ...*exec.Cmd) ([]byte, error) {
	for i, command := range commands[:len(commands)-1] {
		out, err := command.StdoutPipe()
		if err != nil {
			return nil, err
		}
		command.Start()
		commands[i+1].Stdin = out
	}
	final, err := commands[len(commands)-1].Output()
	if err != nil {
		return nil, err
	}
	return final, nil
}
