package migration

import(
	"github.com/kai-zenn/bljr_go_api/api/model"
	"github.com/kai-zenn/bljr_go_api/api/configs"
)

func init(){
	configs.InitDB()
}

func Migrate() {
	configs.DB.AutoMigrate(
		&model.User{},
		&model.Book{},
	)
}