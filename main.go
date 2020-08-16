package main

import (
	// "database/sql"
	// "encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jayzyaj/go-book-store/configs"
	"github.com/jayzyaj/go-book-store/controllers"
	// "github.com/jayzyaj/go-book-store/models"
	"log"
	"net/http"
)

func main() {
	db := configs.InitDB()
	router := mux.NewRouter()
	controllers := controllers.Controllers{}

	router.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controllers.AddBook(db)).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	var book models.Book
// 	params := mux.Vars(r)

// 	err := db.Where("id = ?", params["id"]).First(&book).Error

// 	if err != nil {
// 		panic(err.Error())
// 		return
// 	}

// 	var updatedBook models.Book
// 	json.NewDecoder(r.Body).Decode(&updatedBook)

// 	book.Title = updatedBook.Title
// 	book.Author = updatedBook.Author
// 	book.Year = updatedBook.Year

// 	db.Save(&book)

// 	json.NewEncoder(w).Encode(book)
// }

// func removeBook(w http.ResponseWriter, r *http.Request) {
// 	var book models.Book
// 	params := mux.Vars(r)

// 	err := db.Where("id = ?", params["id"]).First(&book).Error

// 	if err != nil {
// 		panic(err.Error())
// 		return
// 	}

// 	db.Delete(&book)

// 	json.NewEncoder(w).Encode(book)
// }
