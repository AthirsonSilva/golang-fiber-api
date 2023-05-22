package database

import (
	"log"
	"os"

	"github.com/sixfwa/fiber-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {
	log.Println("Database connection starting")

	db, err := gorm.Open(postgres.Open(
		"host=localhost user=postgres password='@Potter77' dbname=go-store port=5432 sslmode=disable"),
		&gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Database connection successfully opened")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Database migration started")

	db.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{})

	Database = DbInstance{Db: db}
}
