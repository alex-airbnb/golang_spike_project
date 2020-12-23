package main

import (
	"log"

	"github.com/alex-airbnb/golang_spike_project/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.SetUp()

}
