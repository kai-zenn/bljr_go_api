package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Book struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key"` 
	Title string `json:"title"`
	Year int `json:"year"`
	Author string `json:"author" gorm:"-"`
	UserID uuid.UUID `json:"-" gorm:"type:uuid;not null"`
	User User `json:"-" gorm:"foreign_key:UserID;references:ID"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}