package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() { // fungsi untuk melakukan koneksi ke database
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_tokoonlen?charset=utf8mb4&parseTime=True&loc=Local")
	//d, err := gorm.Open("mysql", "rudychandra:Rudy_chandra0/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) // digunakan untuk menampilkan stact trace error dan mengheentikan flow goroutine
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
