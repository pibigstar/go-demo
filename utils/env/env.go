package env

import "os"

func IsCI() bool {
	name, _ := os.Hostname()
	if name == "pibigstar" {
		return false
	}
	return true
}
