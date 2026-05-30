package model

type Role struct {
  RoleId uint `json:"id" gorm:"primaryKey;column:role_id"`
  RoleName string `json:"role_name"`
  Accesses []Access `gorm:"many2many:role_access;foreignKey:RoleId;joinForeignKey:RoleId;references:AccessId;joinReferences:AccessId"`
  Users []User `gorm:"many2many:user_role;foreignKey:RoleId;joinForeignKey:role_id;references:ID;joinReferences:user_id"`
}
