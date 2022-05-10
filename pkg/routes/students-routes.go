package routes

import (
	"github.com/Wordyka/go_students_crud_mysql/pkg/controllers"
	"github.com/gorilla/mux"
)

// membuat route yang memanggil controller sebagai routing user mengakses data pada tabel

var RegisterStudentsRoutes = func(router *mux.Router) {
	// Membuat route untuk mengatur URL dan menset method yang dijalankan dari controller pada setiap URL yang diakses
	router.HandleFunc("/produk/", controllers.CreateProduk).Methods("POST")
	router.HandleFunc("/produk/", controllers.GetProduk).Methods("GET")
	router.HandleFunc("/produk/{produkId}", controllers.GetProdukById).Methods("GET")
	router.HandleFunc("/produk/{produkId}", controllers.UpdateProduk).Methods("PUT")
	router.HandleFunc("/produk/{produkId}", controllers.DeleteProduk).Methods("DELETE")

}
