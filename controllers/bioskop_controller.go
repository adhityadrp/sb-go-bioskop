package controllers

import (
	"sb-go-bioskop/config"
	"sb-go-bioskop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var input models.Bioskop

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan lokasi tidak boleh kosong"})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil ditambahkan", "data": input})
}

func GetAllBioskop(c *gin.Context) {
	var bioskops []models.Bioskop
	config.DB.Find(&bioskops)
	c.JSON(http.StatusOK, gin.H{"data": bioskops})
}
