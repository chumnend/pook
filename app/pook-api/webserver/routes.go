package webserver

import (
	"github.com/chumnend/pook/app/pook-api/routes/album"
	"github.com/chumnend/pook/app/pook-api/routes/ping"
	"github.com/gin-gonic/gin"
)

func MakeRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", ping.Pong)

	router.GET("/albums", album.GetAll)
	router.GET("/albums/:id", album.GetByID)
	router.POST("/albums", album.Post)

	return router
}
