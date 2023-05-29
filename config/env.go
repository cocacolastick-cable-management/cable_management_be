package config

var (
	ENV env
)

//mailHost string, host string, port string, password string

type env struct {
	DbDsn string `env:"DB_DSN"`

	JwtSecret string `env:"JWT_SECRET"`
	JwtAuthor string `env:"JWT_AUTHOR"`

	SmtpEmail    string `env:"SMTP_EMAIL"`
	SmtpHost     string `env:"SMTP_HOST"`
	SmtpPort     string `env:"SMTP_PORT"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
}

func init() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//ENV.DbDsn = os.Getenv("DB_DSN")
	ENV.DbDsn = "host=localhost user=postgres password=postgrespw dbname=cable_management_db port=32768 sslmode=disable TimeZone=Asia/Shanghai"

	ENV.JwtSecret = "124567890!@#$%^&*()sdghjklWERTYUIO"
	ENV.JwtAuthor = "vudeptrai"

	ENV.SmtpEmail = "vuphamlethanh@gmail.com"
	ENV.SmtpHost = "smtp.gmail.com"
	ENV.SmtpPort = "587"
	ENV.SmtpPassword = "fzjugwhesxnbpixp"
}
