package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/writecodex09/mahasiswaapicrud/models"
)

type MahasiswaInput struct {
	Nim string 	`json:"nim" binding:"min=5"`
	Nama string `json:"nama"`
}

//Tampilkan data
func Tampil(c *gin.Context) {
	db:= c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(200, gin.H{
		"data" : mhs,
	})
}

//Tambah Data Mahasiswa
func Tambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

//	validasi data -> JSON
var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(400, gin.H{
			"error" : err.Error(),
		})
		return
	}
//	proses input data
mhs := models.Mahasiswa{
	Nim:  dataInput.Nim,
	Nama: dataInput.Nama,
	}
	db.Create(&mhs)
c.JSON(200, gin.H{
	"Messaage" : "Data berhasil di input",
	"Data" : mhs,
})
}
//Ubah Data Mahasiswa
func Ubah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data nya apakah ada atau tidak
	var mhs models.Mahasiswa
	if err:= db.Where("nim = ? ", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(400, gin.H{"error":"data mahasiswa tdak di temukan"})
	}

	//	validasi data -> JSON
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput);err != nil {
		c.JSON(400, gin.H{
			"error" : err.Error(),
		})
		return
	}
	//db.Create(&mhs)
	db.Model(&mhs).Update(dataInput)
	c.JSON(200, gin.H{
		"Messaage" : "Data berhasil di ubah",
		"Data" : mhs,
	})
}
//Hapus Data Mahasiswa
func Hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data nya apakah ada atau tidak
	var mhs models.Mahasiswa
	if err:= db.Where("nim = ? ", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(400, gin.H{"error":"data mahasiswa tdak di temukan"})
	}
	//db.Create(&mhs)
	db.Delete(&mhs)
	c.JSON(200, gin.H{
		"Messaage" : "Data berhasil di Hapus",
		"Data" : true,
	})
}

