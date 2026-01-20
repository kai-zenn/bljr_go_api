package routes

import (
	"github.com/kai-zenn/bljr_go_api/api/controller"
	"github.com/gin-gonic/gin"
)

func BookRoute(r *gin.Engine) {
	userGroup := r.Group("/books")
	{
		userGroup.POST("", controller.CreateBook)
		userGroup.GET("/:id", controller.GetBook)
		userGroup.GET("", controller.GetAllBooks)
		userGroup.PUT("/:id", controller.UpdateBook)
		userGroup.DELETE("/:id", controller.DeleteBook)
	}
}