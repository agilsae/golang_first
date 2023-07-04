package main

import (
	"diary-api/database"
	"diary-api/model"
	"diary-api/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
    loadEnv()
    loadDatabase()
    router.ServeApplication()
}

func loadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&model.User{})
    database.Database.AutoMigrate(&model.Entry{})
}


