package controllers

import (
	"books/database"
	"books/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// *parameter c yang merupakan objek dari gin.Context.
func CreateBook(c *gin.Context) {
	//koneksi ke database
	database.StartDB()

	//mendapatkan instance database yg digunakan pada db.go
	var db = database.GetDB()

	//variable input dengan tipe data models.Book
	var input models.Book

	//mengambil input dari request HTTP yang berupa JSON
	//ShouldBindJson -> parsing body menjadi JSON -> simpan di &input
	if err := c.ShouldBindJSON(&input); err != nil {
		//gin.H{} adalah tipe data map untuk menyimpan data yang dikirimkan sebagai response
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// inisialisasi variabel bookInput dengan nilai dari struct models.Book
	// Struct models.Book adalah representasi dari tabel books dalam database
	bookInput := models.Book{NameBook: input.NameBook, Author: input.Author}
	//method Create pada instance database untuk menyimpan objek
	db.Create(&bookInput)

	c.JSON(http.StatusOK, gin.H{
		"data": bookInput})
}

func GetAllBooks(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	//slice models.Book untuk hasil data
	var book []models.Book

	//mengambil semua data dari tabel books dalam database dan dimasukkan ke dalam slice book
	err := db.Find(&book).Error

	if err != nil {
		fmt.Println("Error getting book data:", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetOneBook(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var bookOne []models.Book

	//First -> untuk mengambil data buku dari database berdasarkan id
	//"Id = ?" -> kriteria pencarian dari parameter c.Param("id")
	err := db.First(&bookOne, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data one": bookOne})
}

func UpdateBook(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var book []models.Book

	err := db.First(&book, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	//input dari request body
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//variable tampung hasil update
	bookUpdate := models.Book{NameBook: input.NameBook, Author: input.Author}
	//Model -> menentukan model yg akan diupdate
	//Where -> kondisi pencarian kolom id dengan nilai param ("id")
	//Updates -> pembaruan record
	db.Model(&input).Where("Id = ?", c.Param("id")).Updates(&bookUpdate)

	c.JSON(http.StatusOK, gin.H{"data": bookUpdate})
}

func DeleteBook(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var bookDelete models.Book
	err := db.First(&bookDelete, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	//Delete -> delete record
	db.Delete(&bookDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
