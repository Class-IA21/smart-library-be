package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvDatabase struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func GetEnvDatabase() *EnvDatabase {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &EnvDatabase{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
}
