package handler

import (
	"fmt"
	"net/http"

	// "log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/book"
)

type bookHanlder struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHanlder {
	return &bookHanlder{bookService}
}

func (handler *bookHanlder) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Mohamad Asep Saepulloh",
		"bio":  "A junior software development",
	})
}

func (handler *bookHanlder) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content":  "Hello World",
		"subtitle": "Belajar Golang 3 jam",
	})
}

func (handler *bookHanlder) BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (handler *bookHanlder) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func (handler *bookHanlder) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		// log.Fatal(err) // ketika error langsung exit terminal
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := handler.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
		// "sub_title": bookRequest.SubTitle,
	})
}
