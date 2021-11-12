package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDB - starts up a new db, with our docker image.
func NewDB() (*gorm.DB, error) {

	fmt.Println("setting up new db connection")
	// 1. bring in credentials to open our stuff.
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	// takes a format layout, and returns it as a string.
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	postgresDB, err := gorm.Open("postgres", connStr)
	if err != nil {
		return postgresDB, err
	}

	if err := postgresDB.DB().Ping(); err != nil {
		return postgresDB, err
	}
	return postgresDB, nil
}
