package utils

import (
	"testing"
)

func TestBase64(t *testing.T) {
	str := "hello world"
	encryptStr := encrypt(str)
	unEncrypt := UnEncrypt(encryptStr)
	if str != unEncrypt {
		t.Error("Failed to encrypt")
	}
}
