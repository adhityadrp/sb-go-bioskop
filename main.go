package main

import (
	"log"
	"os"
	"sb-go-bioskop/config"
	"sb-go-bioskop/models"
	"sb-go-bioskop/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Bioskop{})

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
