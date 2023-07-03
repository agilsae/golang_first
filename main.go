package main

import (
	"diary-api/database"
	"diary-api/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
    loadEnv()
    loadDatabase()
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&model.User{})
    database.Database.AutoMigrate(&model.Entry{})
}

func loadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}
