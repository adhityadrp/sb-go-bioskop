package controllers

import (
	"sb-go-bioskop/config"
	"sb-go-bioskop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CREATE BY ID
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

// GET ALL
func GetAllBioskop(c *gin.Context) {
	var bioskop []models.Bioskop
	result := config.DB.Find(&bioskop)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, bioskop)
}

// GET BY ID
func GetBioskopByID(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}

// UPDATE
func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan lokasi wajib diisi"})
		return
	}

	bioskop.Nama = input.Nama
	bioskop.Lokasi = input.Lokasi
	bioskop.Rating = input.Rating

	config.DB.Save(&bioskop)
	c.JSON(http.StatusOK, gin.H{"message": "Data bioskop berhasil diperbarui", "data": bioskop})
}

// DELETE
func DeleteBioskop(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	config.DB.Delete(&bioskop)
	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}
