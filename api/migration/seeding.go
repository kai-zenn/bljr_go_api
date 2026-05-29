package migration

import (
	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/model"
)


func SeedingDB() {
  var countAccess int64
  configs.DB.Model(&model.Access{}).Count(&countAccess)
  if countAccess == 0 {
    accesses := []model.Access{
      {AccessId: 1, AccessName: "create"},
			{AccessId: 2, AccessName: "read"},
			{AccessId: 3, AccessName: "update"},
			{AccessId: 4, AccessName: "delete"},
    }
    configs.DB.Create(&accesses)
  }

  var countRole int64
  configs.DB.Model(&model.Role{}).Count(&countRole)
  if countRole == 0 {
    roles := []model.Role{
			{RoleId: 1, RoleName: "admin"},
			{RoleId: 2, RoleName: "member"},
			{RoleId: 3, RoleName: "penulis"},
		}
		configs.DB.Create(&roles)
  }
}
