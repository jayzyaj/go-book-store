package configs

import (
	"fmt"
	"github.com/jayzyaj/go-book-store/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// InitDB ...
func InitDB() *gorm.DB {
	db, _ = gorm.Open("mysql", "root:79056123@/books?parseTime=true")

	db.CreateTable(&models.Book{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Book{})
	db.AutoMigrate(&models.Book{})

	fmt.Println("DB Initialized")

	return db
}
