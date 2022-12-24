package app

import (
	"net/http"

	"github.com/chumnend/pook/config"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run(":" + cfg.Port)
}
