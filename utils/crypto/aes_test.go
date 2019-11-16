package utils

import (
	"encoding/hex"
	"testing"
)

var plaintext = "Hello pibigstar"
var key = "11111111111111111111111111111111"
var iv = "123456789"

func TestAesCbc(t *testing.T) {
	var r, _ = AesCbcEncrypt([]byte(plaintext), []byte(key), []byte(iv))
	t.Log("AES CBC Encrypt: ", hex.EncodeToString(r))

	r, _ = AesCbcDecrypt(r, []byte(key), []byte(iv))
	t.Log("AES CBC Decrypt: ", string(r))
}

func TestAesCfb(t *testing.T) {
	encryptStr, _ := AesCfbEncrypt([]byte(plaintext), []byte(key), []byte(iv))
	t.Log("AES CFB Encrypt: ", hex.EncodeToString(encryptStr))

	decryptStr, _ := AesCfbDecrypt(encryptStr, []byte(key), []byte(iv))
	t.Log("AES CFB Decrypt: ", string(decryptStr))
}
