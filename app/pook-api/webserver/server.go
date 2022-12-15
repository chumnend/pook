package webserver

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Name   string
	Router *gin.Engine
	Config *Config
}

func New() (*Server, error) {
	// load config
	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}

	// setup router
	router := MakeRouter()

	s := &Server{
		Name:   "Pook",
		Router: router,
		Config: cfg,
	}

	return s, nil
}

func (s *Server) Start() {
	s.Router.Run(":" + s.Config.Port)
}
