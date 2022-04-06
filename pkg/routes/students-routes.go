package routes
import (
	"github.com/gorilla/mux"
	"github.com/Wordyka/go_students_crud_mysql/pkg/controllers"
)

// membuat route yang memanggil controller sebagai routing user mengakses data pada tabel

var RegisterStudentsRoutes = func(router *mux.Router) {
	router.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")	// routing create
	router.HandleFunc("/student/", controllers.GetStudent).Methods("GET")	// routing get all
	router.HandleFunc("/student/{studentId}", controllers.GetStudentById).Methods("GET")	// routing get by nim
	router.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")		// routing update
	router.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE")	// routing delete
}