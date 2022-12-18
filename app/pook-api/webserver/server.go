package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
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

	// connect database
	db, err := gorm.Open("postgres", cfg.DB)
	if err != nil {
		return nil, err
	}

	// setup router
	router := MakeRouter(cfg, db)

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
