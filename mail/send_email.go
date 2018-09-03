package mail

import (
	"log"
	"net/smtp"
)

const (
	username = "741047261@qq.com"
	password = "授权码"
	host     = "smtp.qq.com"
	addr     = "smtp.qq.com:25"
)

// 发送邮件
func SendEmail() {

	auth := smtp.PlainAuth("", username, password, host)

	user := "741047261@qq.com"
	to := []string{"123456789@qq.com"}
	msg := []byte("this is a message")

	err := smtp.SendMail(addr, auth, user, to, msg)
	if err != nil {
		log.Printf("Error when send email:%s", err.Error())
	}
}
