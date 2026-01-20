package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Year int `json:"year"`
	Author string `json:"author"`
	UserID int `json:"user_id"`
}