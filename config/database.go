package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Get environment variables with fallback values
	dbHost := os.Getenv("PGHOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbUser := os.Getenv("PGUSER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("PGPASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}

	dbName := os.Getenv("PGDATABASE")
	if dbName == "" {
		dbName = "sanbercode_golang_bioskop"
	}

	dbPort := os.Getenv("PGPORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	// Check if DATABASE_URL is provided (Railway specific)
	dbUrl := os.Getenv("DATABASE_URL")
	var dsn string

	if dbUrl != "" {
		// If DATABASE_URL is provided, use it directly
		dsn = dbUrl
	} else {
		// Otherwise, construct the connection string from individual parameters
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPassword, dbName, dbPort)
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi ke database")
	}
	DB = database
}
