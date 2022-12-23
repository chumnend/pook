package webserver

import (
	"github.com/chumnend/pook/app/pook-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func MakeRouter(cfg *Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	api := router.Group("/v1")
	{
		api.GET("/ping", controllers.Ping)

		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
	}

	return router
}
