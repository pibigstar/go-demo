package env

import (
	"os"
	"runtime"
)

func IsCI() bool {
	name, _ := os.Hostname()
	if name != "pibigstar" && runtime.GOOS == "linux" {
		return true
	}
	return false
}
