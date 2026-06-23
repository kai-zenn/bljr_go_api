package model

type Role struct {
  RoleId uint `json:"id" gorm:"primaryKey;column:role_id"`
  RoleName string `json:"role_name" gorm:"column:role_name;not null;unique"`
  Accesses []Access `json:"-" gorm:"many2many:role_access;foreignKey:RoleId;joinForeignKey:RoleId;references:AccessId;joinReferences:AccessId"`
  Users []User `json:"-" gorm:"many2many:user_role;foreignKey:RoleId;joinForeignKey:role_id;references:ID;joinReferences:user_id"`
}
