package config

import (
	"fmt"
	"os"

	"go-api/entities"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbUser,
		dbPass,
		dbDatabase,
		dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entities.Document{}, &entities.User{})

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}