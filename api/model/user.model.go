package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uuid.UUID `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name" gorm:"not null;unique"`
	LastName string `json:"last_name" gorm:"not null;unique"`
	Username string `json:"username" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique;not null"` 
	Password string `json:"-" gorm:"not null"`
	Books []Book `json:"books" gorm:"foreign_key:UserID`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}