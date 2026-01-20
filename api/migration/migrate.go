package main

import(
	"bljr_go_api/api/model"
	"bljr_go_api/api/configs"
)

func init(){
	configs.InitDB()
}

func main() {
	configs.DB.AutoMigrate(
		&model.Book{},
	)
}