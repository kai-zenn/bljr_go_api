package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kai-zenn/bljr_go_api/api/controller"
	"github.com/kai-zenn/bljr_go_api/api/middlewares"
)

func BookRoute(r *gin.Engine) {
	userGroup := r.Group("/books")
	{
		userGroup.POST("", middlewares.AuthMiddleware() ,controller.CreateBook)
		userGroup.GET("/:id", middlewares.AuthMiddleware() , controller.GetBook)
		userGroup.GET("", middlewares.AuthMiddleware() , controller.GetAllBooks)
		userGroup.PUT("/:id",  middlewares.AuthMiddleware() ,controller.UpdateBook)
		userGroup.DELETE("/:id",  middlewares.AuthMiddleware() ,controller.DeleteBook)
	}
}
