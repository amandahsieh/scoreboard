package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DB_USER		string
	DB_PASSWORD	string
	DB_HOST		string
	DB_PORT		string
	DB_NAME		string
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	if DB_USER=="" || DB_PASSWORD=="" || DB_HOST=="" || DB_PORT=="" || DB_NAME=="" {
		log.Fatal("Database configuration missing!")
	}
}