package services

import (
	"fmt"
	emailCont "github.com/cable_management/cable_management_be/src/domain/contracts/email"
)

type IEmailService interface {
	SendEmailNewUser(role string, email string, password string) error
}

type EmailService struct {
	emailHelper emailCont.IEmailHelper
}

func NewEmailService(emailHelper emailCont.IEmailHelper) *EmailService {
	return &EmailService{emailHelper: emailHelper}
}

func (e EmailService) SendEmailNewUser(role string, email string, password string) error {

	err := e.emailHelper.SendEmail(&emailCont.EmailData{
		To:      email,
		Subject: "Your Account",
		Body:    fmt.Sprintf("\n role: %v \n email: %v \n password: %v \n", role, email, password),
	})

	return err
}
