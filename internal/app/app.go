package app

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

func Run(cfg *config.Config) {
	router := gin.Default()

	// connect to database
	pg, err := gorm.Open("postgres", cfg.DB)
	if err != nil {
		log.Fatalf("postgres connect error: %s", err)
	}
	defer pg.Close()

	// setup routes
	v1 := router.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run(":" + cfg.Port)
}
