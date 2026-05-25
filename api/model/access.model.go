package model

import "github.com/google/uuid"


type Access struct {
  AccessId uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
  AccessName string `json:"access_name"`
  Roles []Role `gorm:"many2many:role_access;foreignKey:AccessId;joinForeignKey:AccessId;refrences:RoleId;joinRefrences:RoleId"`
}
