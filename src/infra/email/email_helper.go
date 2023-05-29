package email

import (
	"github.com/cable_management/cable_management_be/src/domain/contracts/email"
	"log"
	"net/smtp"
)

type EmailHelper struct {
	config *email.EmailConfig
}

func NewEmailHelper(config *email.EmailConfig) *EmailHelper {
	return &EmailHelper{config: config}
}

func (eh EmailHelper) SendEmail(data *email.EmailData) error {

	mail := "From: " + eh.config.MailHost + "\n" +
		"To: " + data.To + "\n" +
		"Subject: " + data.Subject + "\n" +
		"\n" +
		data.Body

	auth := smtp.PlainAuth("", eh.config.MailHost, eh.config.Password, eh.config.Host)

	err := smtp.SendMail(eh.config.Host+":"+eh.config.Port, auth, eh.config.MailHost, []string{data.To}, []byte(mail))
	if err != nil {
		log.Fatal(err)
	}

	return err
}
