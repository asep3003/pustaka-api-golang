package main

import (
	"log"

	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	bookRequest := book.BookRequest{
		Title: "Atomic Habbits",
		Price: "150000",
	}

	bookService.Create(bookRequest)

	// bookRepository := book.NewRepository(db)

	// books, err := bookRepository.FindAll()
	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	// book, err := bookRepository.FindById(2)
	// fmt.Println("Title :", book.Title)

	// book := book.Book{
	// 	Title:       "Matahari Minor",
	// 	Description: "Buku series bumi yang ke 4",
	// 	Price:       97000,
	// 	Rating:      8,
	// 	Discount:    5,
	// }
	// bookRepository.Create(book)

	// ===========
	// Create Data
	// ===========

	// book := book.Book{}
	// book.Title = "Bulan"
	// book.Price = 90000
	// book.Discount = 15
	// book.Rating = 8
	// book.Description = "Buku dari serial Bumi Tere Liye yang ke 2"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("=======> Error creating book <=======")
	// }

	// ===========
	// Read Data
	// ===========

	// // var book book.Book
	// var books []book.Book

	// // err = db.Debug().First(&book).Error // ambil data paling atas, limit 1
	// // err = db.Debug().First(&book, 2).Error // ambil data berdasarkan id
	// // err = db.Debug().Find(&books).Error // ambil semua data
	// err = db.Debug().Where("rating = ?", 9).Find(&books).Error
	// if err != nil {
	// 	fmt.Println("=======> Error finding book <=======")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	// ===========
	// Update Data
	// ===========

	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("=======> Error finding book <=======")
	// }

	// book.Title = "Matahari (v2)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("=======> Error updating book <=======")
	// }

	// ===========
	// Delete Data
	// ===========

	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("=======> Error finding book <=======")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("=======> Error deleting book <=======")
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
	// router.Run(":3013") // custom port
}
