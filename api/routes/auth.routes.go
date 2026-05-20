package routes

import (
	"github.com/kai-zenn/bljr_go_api/api/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
  authGroups := r.Group("/auth")
  {
    authGroups.POST("/login", controller.LoginUserHandler)
    authGroups.POST("/register", controller.RegisterUserHandler)
  }
}
