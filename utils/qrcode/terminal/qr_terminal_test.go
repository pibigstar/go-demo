package terminal

import "testing"

func TestGenQrToTerminal(t *testing.T) {
	GenQrToTerminal("http://www.baidu.com")
}

func TestGenQrToTerminalWithConfig(t *testing.T) {
	GenQrToTerminalWithConfig("http://www.baidu.com")
}

func TestGenQR(t *testing.T) {
	GenQR("Hello World")
}
