package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alex-airbnb/golang_spike_project/model"
	"github.com/joho/godotenv"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Article{})
}

// SetUp Set up the connection with the Postgres DB.
func SetUp() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect to the database")
	}

	migrate(db)

	return db
}
