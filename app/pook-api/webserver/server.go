package webserver

import (
	"github.com/gin-gonic/gin"

	"github.com/chumnend/pook/app/pook-api/routes/album"
)

type Server struct {
	name   string
	router *gin.Engine
}

func New() (*Server, error) {
	// setup router
	router := gin.Default()
	router.GET("/albums", album.GetAll)
	router.GET("/albums/:id", album.GetByID)
	router.POST("/albums", album.Post)

	s := &Server{name: "Pook", router: router}

	return s, nil
}

func (s *Server) Start() {
	s.router.Run("localhost:8080")
}
