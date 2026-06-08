package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kai-zenn/bljr_go_api/api/controller"
	"github.com/kai-zenn/bljr_go_api/api/middlewares"
)

func SetupTestRoutes(r *gin.Engine) {
  // ini kode dari tutor jir
  // Public routes (accessible by anyone)
   // r.POST("/login", controller.LoginUserHandler)
   // r.POST("/register", controller.RegisterUserHandler)

   // Protected routes (secured with RBAC)
   authGroup := r.Group("/rbac")
   authGroup.Use(middlewares.AuthMiddleware())

   // Role-based access control for specific routes
   authGroup.GET("/users", middlewares.RBACMiddleware("read"), controller.GetUsers)
   authGroup.GET("/users/:id", middlewares.RBACMiddleware("read"), controller.GetUserById)
   authGroup.PUT("/users/:id", middlewares.RBACMiddleware("update"), controller.UpdateUser)
}
