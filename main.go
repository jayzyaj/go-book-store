package main

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jayzyaj/go-book-store/controllers"
	"github.com/jayzyaj/go-book-store/models"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open("mysql", "root:79056123@/books?parseTime=true")

	db.CreateTable(&models.Book{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Book{})
	db.AutoMigrate(&models.Book{})

	// var newBook = Book{Title: "Gago", Author: "Badang", Year: "2010"}
	// db.Create(&newBook)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// err = db.Ping()

	// if err != nil {
	// 	panic(err.Error())
	// }
	fmt.Println("Connected")

	// defer db.Close()
}

func main() {
	router := mux.NewRouter()
	controllers := controllers.Controllers{}

	router.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	err := db.Where("id = ?", params["id"]).First(&book).Error

	if err != nil {
		panic(err.Error())
		return
	}

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	err := db.Create(&book).Error

	if err != nil {
		panic(err.Error())
		return
	}

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	err := db.Where("id = ?", params["id"]).First(&book).Error

	if err != nil {
		panic(err.Error())
		return
	}

	var updatedBook models.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Year = updatedBook.Year

	db.Save(&book)

	json.NewEncoder(w).Encode(book)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	err := db.Where("id = ?", params["id"]).First(&book).Error

	if err != nil {
		panic(err.Error())
		return
	}

	db.Delete(&book)

	json.NewEncoder(w).Encode(book)
}
