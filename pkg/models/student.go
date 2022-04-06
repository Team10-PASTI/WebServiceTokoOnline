package models

import (
	"github.com/Wordyka/go_students_crud_mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// mendefenisikan suatu struct dengan nama Student sebagai field dari tabel database
type Student struct {
	gorm.Model
	NIM           string `gorm:""json:"nim"`
	Name          string `json:"name"`
	IPK           string `json:"ipk"`
	Jurusan       string `json:"jurusan"`
	Angkatan      string `json:"angkatan"`
	StatusAktif   string `jsonn:"status_aktif"`
	Username      string `json:"username"`
	EmailAkademik string `json:"email_akademik"`
	WaliMahasiswa string `json:"wali_mahasiswa"`
	JalurUSM      string `json:"jalur_USM"`
}

// fungsi init untuk  menginisialisasi koneksi kepada database
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}


// fungsi untuk membuat data baru pada tabel student
func (b *Student) CreateStudent() *Student {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// fungsi untuk mengambil semua data pada tabel student
func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}

// fungsi untuk mengambil semua data pada tabel student berdasarkan nim
func GetStudentbyId(nim int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("NIM=?", nim).Find(&getStudent)
	return &getStudent, db
}

// fungsi untuk menghapus salah satu data pada tabel student berdasarkan nim
func DeleteStudent(nim int64) Student {
	var student Student
	db.Where("NIM=?", nim).Delete(student)
	return student
}
