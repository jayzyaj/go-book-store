package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}
