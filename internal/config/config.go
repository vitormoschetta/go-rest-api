package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if os.Getenv("ENV") != "production" {
		loadLocalEnv()
	}
}

func loadLocalEnv() {
	log.Println("Loading local .env file...")

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	log.Println("DATABASE_HOST: " + os.Getenv("DATABASE_HOST"))
	log.Println("DATABASE_PORT: " + os.Getenv("DATABASE_PORT"))
	log.Println("DATABASE_NAME: " + os.Getenv("DATABASE_NAME"))
	log.Println("DATABASE_USER: " + os.Getenv("DATABASE_USER"))
	log.Println("DATABASE_PASSWORD: " + os.Getenv("DATABASE_PASSWORD"))
	log.Println("APPLICATION_PORT: " + os.Getenv("APPLICATION_PORT"))

	log.Println("Local .env file loaded!")
}
