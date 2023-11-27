package book

import "encoding/json"

type BookRequest struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
	// Price int    `json:"price" binding:"required,number"` // default tipe integer
	// SubTitle string `json:"sub_title"` // set untuk field usable
}
