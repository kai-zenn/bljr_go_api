package main

import (
	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/routes"
	"github.com/kai-zenn/bljr_go_api/api/migration"
	"github.com/gin-gonic/gin"
)

func Init() {
	configs.InitDB()
	migration.Migrate()
}

func main() {

	Init()
	r := gin.Default()

	routes.BookRoute(r)
	routes.UserRoute(r)
	routes.AuthRoutes(r)

	r.Run()
}
