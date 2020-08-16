package controllers

import (
	"encoding/json"
	"github.com/jayzyaj/go-book-store/models"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Controllers struct{}

func (c Controllers) GetBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var books = []models.Book{}

		errors := db.Find(&books).GetErrors()

		for _, err := range errors {
			panic(err.Error())
			return
		}

		json.NewEncoder(w).Encode(books)
	}
}
