package email

type EmailConfig struct {
	MailHost string
	Host     string
	Port     string
	Password string
}

func NewEmailConfig(mailHost string, host string, port string, password string) *EmailConfig {
	return &EmailConfig{MailHost: mailHost, Host: host, Port: port, Password: password}
}

type IEmailHelper interface {
	SendEmail(data *EmailData) error
}
