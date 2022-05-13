package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Wordyka/go_students_crud_mysql/pkg/models"
	"github.com/Wordyka/go_students_crud_mysql/pkg/utils"
	"github.com/gorilla/mux"
)

var NewProduk models.Produk

// fungsi yang digunakan memanggil seluruh data yang terdapat pada tabel
func GetProduk(w http.ResponseWriter, r *http.Request) {
	newProduk := models.GetAllProduk()
	res, _ := json.Marshal(newProduk)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan memanggil data yang sesuai dengan ID yang direquest
func GetProdukById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProdukId := vars["ProdukID"]
	ID, err := strconv.ParseInt(ProdukId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	ProdukDetails, _ := models.GetProdukbyId(ID)
	res, _ := json.Marshal(ProdukDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan ketika menmbah user baru pada database
func CreateProduk(w http.ResponseWriter, r *http.Request) {
	CreateProduk := &models.Produk{}
	utils.ParseBody(r, CreateProduk)
	b := CreateProduk.CreateProduk()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan ketika menhapus data pada database
func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProdukID := vars["ProdukID"]
	ID, err := strconv.ParseInt(ProdukID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	Produk := models.DeleteProduk(ID)
	res, _ := json.Marshal(Produk)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi unutk mengedit data pada database
func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	var updateProduk = &models.Produk{}
	utils.ParseBody(r, updateProduk)
	vars := mux.Vars(r)
	ProdukID := vars["ProdukID"]
	ID, err := strconv.ParseInt(ProdukID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	ProdukDetails, db := models.GetProdukbyId(ID)
	if updateProduk.Nama != "" { // pengkondisian untuk mengubah nilai dari ipk jika terdapat perubahan yang dilakukan
		ProdukDetails.Nama = updateProduk.Nama
	}
	if updateProduk.Kategori != "" { // pengkondisian untuk mengubah nilai dari jurusan jika terdapat perubahan yang dilakukan
		ProdukDetails.Kategori = updateProduk.Kategori
	}
	if updateProduk.Harga != 0 { // pengkondisian untuk mengubah nilai dari angkatan jika terdapat perubahan yang dilakukan
		ProdukDetails.Harga = updateProduk.Harga
	}
	if updateProduk.Gambar != "" { // pengkondisian untuk mengubah nilai dari status aktif jika terdapat perubahan yang dilakukan
		ProdukDetails.Gambar = updateProduk.Gambar
	}
	if updateProduk.Detail != "" { // pengkondisian untuk mengubah nilai dari username jika terdapat perubahan yang dilakukan
		ProdukDetails.Detail = updateProduk.Detail
	}
	if updateProduk.Jumlah != 0 { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		ProdukDetails.Jumlah = updateProduk.Jumlah
	}
	db.Save(&ProdukDetails) // mensave hasil perubahan
	res, _ := json.Marshal(ProdukDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
