package models

import (
	"github.com/jinzhu/gorm"
	"github.com/leomuko/go-mysql/pkg/config"
)

var db *gorm.DB

//will need to cross check this
// type Book struct {
// 	gorm.Model
// 	BookId string `gorm:""json:"book_id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// 	Price  int    `json:"price"`
// }

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var theBook Book
	db := db.Where("ID=?", Id).Find(&theBook)
	return &theBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
