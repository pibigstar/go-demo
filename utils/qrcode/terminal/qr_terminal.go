package terminal

import (
	"github.com/mattn/go-colorable"
	"github.com/mdp/qrterminal/v3"
	"os"
	"runtime"
)

// 将二维码输出到控制台
func GenQrToTerminal(content string) {
	qrterminal.Generate(content, qrterminal.L, os.Stdout)
}

// 使用简单配置
func GenQrToTerminalWithConfig(content string) {
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig(content, config)
}

// 复杂点配置
func GenQR(content string) {
	qrConfig := qrterminal.Config{
		HalfBlocks:     true,
		Level:          qrterminal.L,
		Writer:         os.Stdout,
		BlackWhiteChar: "\u001b[37m\u001b[40m\u2584\u001b[0m",
		BlackChar:      "\u001b[30m\u001b[40m\u2588\u001b[0m",
		WhiteBlackChar: "\u001b[30m\u001b[47m\u2585\u001b[0m",
		WhiteChar:      "\u001b[37m\u001b[47m\u2588\u001b[0m",
	}
	if runtime.GOOS == "windows" {
		qrConfig.HalfBlocks = false
		qrConfig.Writer = colorable.NewColorableStdout()
		qrConfig.BlackChar = qrterminal.BLACK
		qrConfig.WhiteChar = qrterminal.WHITE
	}

	qrterminal.GenerateWithConfig(content, qrConfig)
}
