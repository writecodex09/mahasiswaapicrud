package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/korneliusdb?charset=utf8&parseTime=True&loc=Local")
	{
		if err != nil {
			panic("Gagal koneksi database")
		}
		db.AutoMigrate(&Mahasiswa{})
		return db
	}
}


