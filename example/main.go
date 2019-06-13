package main

import (
	"log"

	"github.com/zs5460/mail"
)

func main() {
	cfg := mail.Config{
		MailSubject:   "test",
		MailServer:    "smtp.xxx.com:25",
		MailSender:    "xxx@xxx.com",
		MailSenderPwd: "******",
		MailReciver:   "abc@abc.com",
	}

	err := mail.SendMail(
		cfg.MailSender,
		cfg.MailSenderPwd,
		cfg.MailServer,
		cfg.MailReciver,
		cfg.MailSubject,
		"this is a test mail used:"+cfg.MailServer,
	)
	if err != nil {
		log.Println(err)
	}

}
