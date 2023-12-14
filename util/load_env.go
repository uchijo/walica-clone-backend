package util

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env file not found. going to load from bare env vars.")
	}
}
