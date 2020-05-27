package code

import (
	"github.com/dchest/captcha"
	"os"
)

func GenCodeImage(length int) error {
	f, err := os.Create("code.png")
	if err != nil {
		return err
	}

	return captcha.WriteImage(f, captcha.NewLen(length), length*25, length*10)
}
