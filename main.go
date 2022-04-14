package main

import (
	"github.com/gin-gonic/gin"
	"github.com/writecodex09/mahasiswaapicrud/controller"
	"github.com/writecodex09/mahasiswaapicrud/models"
)

func main() {

	router := gin.Default()
	//models
	db := models.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db",db)
		c.Next()
	})
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"Message" : "Home Page",
		})
	})

	router.GET("/mahasiswa", controller.Tampil)
	router.POST("/mahasiswa", controller.Tambah)
	router.PUT("/mahasiswa/:nim", controller.Ubah)
	router.DELETE("/mahasiswa/:nim", controller.Hapus)
	router.Run()
}
