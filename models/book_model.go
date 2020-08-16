package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	ID     int
	Title  string
	Author string
	Year   string
}
