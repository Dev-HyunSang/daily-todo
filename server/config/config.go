package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to Load .env file")
	}

	return os.Getenv(name)
}
