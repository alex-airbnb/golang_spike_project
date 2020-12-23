package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alex-airbnb/golang_spike_project/config"
	"github.com/alex-airbnb/golang_spike_project/model"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Article{})
}

// SetUp Set up the connection with the Postgres DB.
func SetUp() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect to the database")
	}

	migrate(db)

	return db
}
