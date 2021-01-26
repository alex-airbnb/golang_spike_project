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

// PostgresDB Instance of the Postgres DB
var PostgresDB *gorm.DB

// SetUpPostgres Set up the connection with the Postgres DB.
func SetUpPostgres() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	PostgresDB, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect to the database")

		return err
	}

	log.Print("Set up PostgresDB")

	migrate(PostgresDB)

	return nil
}
