package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jayzyaj/go-book-store/models"
	"github.com/jayzyaj/go-book-store/repository"
	"github.com/jayzyaj/go-book-store/utils"
	"github.com/jinzhu/gorm"
	"net/http"
)

// Controllers ...
type Controllers struct{}

// GetBooks ...
func (c Controllers) GetBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var books = []models.Book{}
		var error models.Error
		bookRepo := repository.BookRepository{}

		books, err := bookRepo.GetBooks(db, books)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

// GetBook ...
func (c Controllers) GetBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)
		bookRepo := repository.BookRepository{}

		book, err := bookRepo.GetBook(db, book, params["id"])

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

// AddBook ...
func (c Controllers) AddBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)
		bookRepo := repository.BookRepository{}

		book, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

// UpdateBook ...
func (c Controllers) UpdateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)
		bookRepo := repository.BookRepository{}

		var newBook models.Book
		json.NewDecoder(r.Body).Decode(&newBook)

		book, err := bookRepo.UpdateBook(db, book, newBook, params["id"])

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

// DeleteBook ...
func (c Controllers) DeleteBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)
		bookRepo := repository.BookRepository{}

		book, err := bookRepo.Deletebook(db, book, params["id"])

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}
