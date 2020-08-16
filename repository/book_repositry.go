package repository

import (
	"github.com/jayzyaj/go-book-store/models"
	"github.com/jinzhu/gorm"
)

// BookRepository ...
type BookRepository struct{}

// GetBooks ...
func (b BookRepository) GetBooks(db *gorm.DB, books []models.Book) ([]models.Book, error) {
	err := db.Find(&books).Error
	if err != nil {
		return []models.Book{}, err
	}
	return books, err
}

// GetBook ...
func (b BookRepository) GetBook(db *gorm.DB, book models.Book, id string) (models.Book, error) {
	err := db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return models.Book{}, err
	}
	return book, err
}

// AddBook ...
func (b BookRepository) AddBook(db *gorm.DB, book models.Book) (models.Book, error) {
	err := db.Create(&book).Error

	if err != nil {
		return models.Book{}, err
	}
	return book, err
}

// UpdateBook ...
func (b BookRepository) UpdateBook(db *gorm.DB, book models.Book, newBook models.Book, id string) (models.Book, error) {
	err := db.Where("id = ?", id).First(&book).Error

	if err != nil {
		return models.Book{}, err
	}

	book.Title = newBook.Title
	book.Author = newBook.Author
	book.Year = newBook.Year

	db.Save(&book)

	return book, err
}

// Deletebook ...
func (b BookRepository) Deletebook(db *gorm.DB, book models.Book, id string) (models.Book, error) {
	err := db.Where("id = ?", id).First(&book).Error

	if err != nil {
		return models.Book{}, err
	}

	error := db.Delete(&book).Error

	if error != nil {
		return models.Book{}, error
	}

	return book, error
}

// err := db.Where("id = ?", params["id"]).First(&book).Error

// 		if err != nil {
// 			panic(err.Error())
// 			return
// 		}

// 		var updatedBook models.Book
// 		json.NewDecoder(r.Body).Decode(&updatedBook)

// 		book = updatedBook

// 		db.Save(&book)
