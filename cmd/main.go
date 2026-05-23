package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/migration"
	"github.com/kai-zenn/bljr_go_api/api/routes"
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

	r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello from Gin!",
        })
    })

	fmt.Println("\n  ➜  Local: http://localhost:6000/")
	r.Run()
}
