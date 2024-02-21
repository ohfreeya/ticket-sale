package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfigStruct struct {
	Host     string
	Username string
	Password string
	Database string
}

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DBConfig() DBConfigStruct {
	return DBConfigStruct{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}
