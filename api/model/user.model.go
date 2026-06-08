package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uuid.UUID `json:"id" gorm:"tyep:char(36);primary_key"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName string `json:"last_name" gorm:"not null"`
	Username string `json:"username" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique;not null"` 
	Password string `json:"-" gorm:"not null"`
	Books []Book `json:"books" gorm:"foreign_key:UserID"`
	Roles []Role `json:"role" gorm:"many2many:user_role;foreignKey:ID;joinForeignKey:user_id;references:RoleID;joinReferences:role_id;"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
