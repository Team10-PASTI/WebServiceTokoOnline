package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// deklarasi variabel melalui package yang diimport sebagai penghubung interaksi ke database
var (
	db *gorm.DB
)

// fungsi untuk melakukan koneksi ke jenis database mysql dengan menginsert username, password, dan tabel database yang akan dikoneksikan 
func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_go_mysql?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// fungsi untuk mendapatkan database 
func GetDB() *gorm.DB {
	return db
}
