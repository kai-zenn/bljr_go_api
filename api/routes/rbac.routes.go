package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kai-zenn/bljr_go_api/api/controller"
	"github.com/kai-zenn/bljr_go_api/api/middlewares"
)

func RbacRoutes(r *gin.Engine) {
  // ini kode dari tutor jir
  // Public routes (accessible by anyone)
   // r.POST("/login", controller.LoginUserHandler)
   // r.POST("/register", controller.RegisterUserHandler)

   // Protected routes (secured with RBAC)
   authGroup := r.Group("/rbac")
   authGroup.Use(middlewares.AuthMiddleware())

   // Role-based access control for routes User Routes (only admin can do CRUD Operations on user)
   authGroup.GET("/users", middlewares.RBACMiddleware("users:read"), controller.GetUsers)
   authGroup.GET("/users/:id", middlewares.RBACMiddleware("users:read"), controller.GetUserById)
   authGroup.PATCH("/users/:id/change-role", middlewares.RBACMiddleware("users:update"), controller.UpdateUser)
   // authGroup.DELETE("/users/:id", middlewares.RBACMiddleware("users:delete"), controller.DeleteUser)
   
   // Role-based access control for Books Routes
   authGroup.GET("/books", middlewares.RBACMiddleware("books:read"), controller.GetAllBooks)
   authGroup.GET("/books/:id", middlewares.RBACMiddleware("books:read"), controller.GetBook)
   authGroup.POST("/books", middlewares.RBACMiddleware("books:write"), controller.CreateBook) // only Writer can create books
   authGroup.PATCH("/books/:id", middlewares.RBACMiddleware("books:write"), controller.UpdateBook)
   authGroup.DELETE("/books/:id", middlewares.RBACMiddleware("books:write"), controller.DeleteBook)
}
