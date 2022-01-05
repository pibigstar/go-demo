// +build windows

package disk

import (
	"runtime"
	"testing"
)

func TestGetSystemDisks(t *testing.T) {
	sysType := runtime.GOOS
	if sysType == "windows" {
		GetSystemDisks()
	}
}
