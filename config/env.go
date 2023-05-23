package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Env env
)

type env struct {
	DbDsn string `env:"DB_DSN"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	Env.DbDsn = os.Getenv("DB_DSN")
}
