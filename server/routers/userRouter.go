package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vesarian/cms-cloth/controllers"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetUser)
	router.POST("/register", controllers.UserRegister)
}