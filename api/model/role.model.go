package model

type Role struct {
  RoleId uint `json:"id" gorm:"primaryKey;column:role_id"`
  RoleName string `json:"role_name"`
  Accesses []Access `gorm:"many2many:role_access;foreignKey;RoleId;joinForeignKey:RoleId;refrences:AccessId;joinRefrences:AccessId"`
  Users []User `gorm:"many2many:user_role;foreignKey:RoleId;joinForeignKey:role_id;refrences:ID;joinRefrences:user_id"`
}
