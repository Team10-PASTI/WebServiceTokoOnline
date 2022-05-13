package models

import (
	"github.com/Wordyka/go_students_crud_mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Membuat sebuah struct unutk mencreate tabel Produk dan tipe data setiap kolomnya
type Produk struct {
	gorm.Model
	Nama     string `gorm:""json:"nama"`
	Kategori string `json:"kategori"`
	Harga    uint   `json:"harga"`
	Gambar   string `json:"gambar"`
	Detail   string `json:"detail"`
	Jumlah   uint   `json:"jumlah"`
}

// Fungsi unutk melakukan migrasi pada table dan akan otomatis berjalan ketika di run dan terdapat perubahan pada struktur table
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Produk{})
}

// fungsi unutuk mencreata user baru
func (b *Produk) CreateProduk() *Produk {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//fungsi untuk mengambil semua data yang terdapat pada table
func GetAllProduk() []Produk {
	var Produk []Produk
	db.Find(&Produk)
	return Produk
}

// fungsi untuk mengambil data table sesuai dengan id yang di request
func GetProdukbyId(id int64) (*Produk, *gorm.DB) {
	var getProduk Produk
	db := db.Where("ID=?", id).Find(&getProduk)
	return &getProduk, db
}

// fungsi untuk menghapus data table sesuai dengan id yang di request
func DeleteProduk(id int64) Produk {
	var produk Produk
	db.Where("ID=?", id).Delete(produk)
	return produk

}
