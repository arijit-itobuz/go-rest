package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_ENV string
	PORT    string
}

var AppConfig Config

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("[Error]: Failed to load .env file")
	}

	AppConfig = Config{
		APP_ENV: os.Getenv("APP_ENV"),
		PORT:    os.Getenv("PORT"),
	}

}
