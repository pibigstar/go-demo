package sdk

import (
	"github.com/skip2/go-qrcode"
)

func GenQRCode(content string, level qrcode.RecoveryLevel, size int) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(content, level, size)
	if err != nil {
		return nil, err
	}
	return png, err
}

func GenDefaultQRCode(content string) ([]byte, error) {
	return GenQRCode(content, qrcode.Medium, 256)
}
