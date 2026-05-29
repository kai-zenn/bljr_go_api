package migration

import(
	"github.com/kai-zenn/bljr_go_api/api/model"
	"github.com/kai-zenn/bljr_go_api/api/configs"
)

func Migrate() {
  err := configs.DB.AutoMigrate(
    		&model.User{},
    		&model.Book{},
    		&model.Role{},
    		&model.Access{},
   	)

  if err != nil {
		panic("Gagal melakukan migrasi database: " + err.Error())
	}
}
