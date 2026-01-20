package main

import (
	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	configs.InitDB()
}

func main() {
	Init()
	r := gin.Default()

	routes.BookRoute(r)

	r.Run(":3000")
}