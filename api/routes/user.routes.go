package routes

import (
	"github.com/kai-zenn/bljr_go_api/api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:id", controller.GetUserById)
		userGroup.GET("", controller.GetUsers)
		userGroup.PUT("/:id", controller.UpdateUser)
	}
}
