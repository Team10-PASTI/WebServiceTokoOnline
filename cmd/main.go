package cmd

// import package yang akan digunakan 
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Wordyka/go_students_crud_mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()	// membuat suatu gorilla / mux router baru
	routes.RegisterStudentsRoutes(r)	// memanggil variabel RegisterStudentRoutes pada package routes
	http.Handle("/", r)	 // fungsi handle untuk mendefeinisikan pattern dari route dan sebagai listen / serve HTTP function
	fmt.Print("Starting Server localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

