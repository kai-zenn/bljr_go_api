package main

import(
	"github.com/kai-zenn/bljr_go_api/api/model"
	"github.com/kai-zenn/bljr_go_api/api/configs"
)

func init(){
	configs.InitDB()
}

func main() {
	configs.DB.AutoMigrate(
		&model.Book{},
	)
}