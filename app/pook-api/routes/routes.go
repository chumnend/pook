package routes

import (
	"github.com/chumnend/pook/app/pook-api/controllers"
	"github.com/chumnend/pook/app/pook-api/middlewares"
	"github.com/gin-gonic/gin"
)

func MakeRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/ping", controllers.Ping)
	v1.POST("/register", controllers.Register)
	v1.POST("/login", controllers.Login)

	admin := router.Group("/admin")
	admin.Use(middlewares.JwtAuthMiddleware())

	admin.GET("/user", controllers.CurrentUser)

	return router
}
