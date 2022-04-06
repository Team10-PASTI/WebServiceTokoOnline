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

// menginstansisasi Student dari package models pada variabel NewStudent
var NewStudent models.Student

// fungsi mendapatkan semua data pada tabel student
func GetStudent(w http.ResponseWriter, r *http.Request) {
	newStudents := models.GetAllStudents()             // memanggil fungsi GetAllStudents() pada package models
	res, _ := json.Marshal(newStudents)                // melakukan konversi (marshalling) object menjadi JSON
	w.Header().Set("Content-Type", "pkglication/json") // membuat header dengan parameter berupa pasangan key dan value
	w.WriteHeader(http.StatusOK)                       // status bahwa data berhasil di get
	w.Write(res)                                       // menampilkan data
}

// fungsi mendapatkan data pada tabel student berdasarkan nim
func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                           // mengambil data dari segmen dengan http.Request sebagai parameter
	studentId := vars["studentId"]                // data yang diambil didasarkan pada value StudentId
	NIM, err := strconv.ParseInt(studentId, 0, 0) // melakukan pengecekan error menggunakan fungsi strconv() untuk konversi
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentDetails, _ := models.GetStudentbyId(NIM)    // memanggil fungsi GetStudentbyId dengan parameter value NIM
	res, _ := json.Marshal(studentDetails)             // melakukan konversi (marshalling) object menjadi JSON
	w.Header().Set("Content-Type", "pkglication/json") // membuat header dengan parameter berupa pasangan key dan value
	w.WriteHeader(http.StatusOK)                       // status bahwa data berhasil di get
	w.Write(res)                                       // menampilkan data
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	CreateStudent := &models.Student{}                 // memanggil struct Student
	utils.ParseBody(r, CreateStudent)                  // memanggil fungsi ParseBody() dari package utils untuk parsing body menjadi byte
	b := CreateStudent.CreateStudent()                 // memanggil fungsi CreateStudent() dari package student
	res, _ := json.Marshal(b)                          // melakukan konversi (marshalling) object menjadi JSON
	w.Header().Set("Content-Type", "pkglication/json") // membuat header dengan parameter berupa pasangan key dan value
	w.WriteHeader(http.StatusOK)                       // status bahwa data berhasil di get
	w.Write(res)                                       // menampilkan data
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                          // mengambil data dari segmen dengan http.Request sebagai parameter
	studentId := vars["studentId"]               // data yang diambil didasarkan pada value StudentId
	ID, err := strconv.ParseInt(studentId, 0, 0) // melakukan pengecekan error menggunakan fungsi strconv() untuk konversi
	if err != nil {
		fmt.Println("error while parsing")
	}
	student := models.DeleteStudent(ID)                // memanggil fungsi DeleteStudent dengan parameter value NIM
	res, _ := json.Marshal(student)                    // melakukan konversi (marshalling) object menjadi JSON
	w.Header().Set("Content-Type", "pkglication/json") // membuat header dengan parameter berupa pasangan key dan value
	w.WriteHeader(http.StatusOK)                       // status bahwa data berhasil di get
	w.Write(res)                                       // menampilkan data
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var updateStudent = &models.Student{}        // memanggil struct Student
	utils.ParseBody(r, updateStudent)            // memanggil fungsi ParseBody() dari package utils untuk parsing body menjadi byte
	vars := mux.Vars(r)                          // mengambil data dari segmen dengan http.Request sebagai parameter
	studentId := vars["studentId"]               // data yang diambil didasarkan pada value StudentId
	ID, err := strconv.ParseInt(studentId, 0, 0) // melakukan pengecekan error menggunakan fungsi strconv() untuk konversi
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentDetails, db := models.GetStudentbyId(ID) // memanggil fungsi GetStudentbyId dari package models

	// pengecekan jika field yang akan diupdate adalah data yang tidak kosong
	if updateStudent.Name != "" {
		studentDetails.Name = updateStudent.Name
	}
	if updateStudent.IPK != "" {
		studentDetails.IPK = updateStudent.IPK
	}
	if updateStudent.Jurusan != "" {
		studentDetails.Jurusan = updateStudent.Jurusan
	}
	if updateStudent.Angkatan != "" {
		studentDetails.Angkatan = updateStudent.Angkatan
	}
	if updateStudent.StatusAktif != "" {
		studentDetails.StatusAktif = updateStudent.StatusAktif
	}
	if updateStudent.Username != "" {
		studentDetails.Username = updateStudent.Username
	}
	if updateStudent.EmailAkademik != "" {
		studentDetails.EmailAkademik = updateStudent.EmailAkademik
	}
	if updateStudent.WaliMahasiswa != "" {
		studentDetails.WaliMahasiswa = updateStudent.WaliMahasiswa
	}
	if updateStudent.JalurUSM != "" {
		studentDetails.JalurUSM = updateStudent.JalurUSM
	}

	db.Save(&studentDetails)                           // menyimpan perubahan pada database
	res, _ := json.Marshal(studentDetails)             // melakukan konversi (marshalling) object menjadi JSON
	w.Header().Set("Content-Type", "pkglication/json") // membuat header dengan parameter berupa pasangan key dan value
	w.WriteHeader(http.StatusOK)                       // status bahwa data berhasil di get
	w.Write(res)                                       // menampilkan data
}
