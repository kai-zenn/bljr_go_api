package model

type Access struct {
  AccessId uint `json:"id" gorm:"primaryKey;column:access_id"`
  AccessName string `json:"access_name"`
  Roles []Role `gorm:"many2many:role_access;foreignKey:AccessId;joinForeignKey:AccessId;references:RoleId;joinReferences:RoleId"`
}
