package main

import (
	"sb-go-bioskop/config"
	"sb-go-bioskop/models"
	"sb-go-bioskop/routes"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Bioskop{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
