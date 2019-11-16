package utils

import (
	"testing"
)

func TestBase64(t *testing.T) {
	str := "hello world"
	encryptStr := encrypt(str)
	unEncrypt := UnEncrypt(encryptStr)
	if str != unEncrypt {
		t.Error("Failed to crypto")
	}
}

func TestBase64Decode(t *testing.T) {
	code := Base64Encode("pibigstar")
	t.Log(code)

	decode := Base64Decode(code)
	t.Log(decode)
}
