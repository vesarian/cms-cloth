package routers

import (
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	usersRouter := router.Group("/users")
	UserRoutes(usersRouter)

	categoriesRouter := router.Group("/categories")
	CategoryRoutes(categoriesRouter)

	clothRouter := router.Group("/cloths")
	ClothRoutes(clothRouter)

	return router
}
