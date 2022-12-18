package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func MakeRouter(cfg *Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	api := router.Group("/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		api.POST("/register", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint"})
		})
		api.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "this is the login endpoint"})
		})
	}

	return router
}
