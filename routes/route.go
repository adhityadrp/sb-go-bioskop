package routes

import (
	"sb-go-bioskop/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/bioskop", controllers.CreateBioskop)
	r.GET("/bioskop", controllers.GetAllBioskop)

	return r
}
