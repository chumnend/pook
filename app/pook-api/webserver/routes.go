package webserver

import (
	"github.com/chumnend/pook/app/pook-api/routes/album"
	"github.com/chumnend/pook/app/pook-api/routes/ping"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func MakeRouter(cfg *Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	api := router.Group("/v1")
	{
		api.GET("/ping", ping.Pong)
	}

	router.GET("/albums", album.GetAll)
	router.GET("/albums/:id", album.GetByID)
	router.POST("/albums", album.Post)

	return router
}
