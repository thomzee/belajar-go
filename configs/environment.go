package configs

import (
	"github.com/joho/godotenv"
	"os"
)

func Env() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	env := os.Getenv("APP_ENV")
	err = godotenv.Load(".env." + env)
	if err != nil {
		panic(err)
	}
}
