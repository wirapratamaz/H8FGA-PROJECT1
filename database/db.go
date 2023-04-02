package database

import (
	"books/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "project1"
	db       *gorm.DB
	err      error
)

func StartDB() {
	//konfigurasi ke database
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode= disable", host, user, password, dbname, port)

	// postgress.Open(config) -> mengembalikan pointer ke struct DB yang merepresentasikan koneksi
	// mengirimkan parameter &gorm.Config{} untuk konfigurasi tambahan GORM
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	// debug mode untuk melakukan migrasi otomatis pada tabel Book
	db.Debug().AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	// mengembalikan koneksi database
	return db
}
