package main

// import package yang akan digunakan
import (
	"fmt"
	"log"
	"net/http"

	"github.com/Wordyka/go_students_crud_mysql/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()                                // memanggil method yang digunkan untuk melakukan routing
	routes.RegisterStudentsRoutes(r)                    // mendaftara
	http.Handle("/", r)                                 // mendaftarkan view ke paket http
	fmt.Print("Starting Server localhost:9010")         // mencetak string saat program berhasil terhubung ke server mysql
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // digunakan untuk membuat sekaligus start server baru
}
