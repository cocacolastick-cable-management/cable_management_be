package config

var (
	ENV env
)

type env struct {
	DbDsn string `env:"DB_DSN"`

	JwtSecret string `env:"JWT_SECRET"`
	JwtAuthor string `env:"JWT_AUTHOR"`
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
}
