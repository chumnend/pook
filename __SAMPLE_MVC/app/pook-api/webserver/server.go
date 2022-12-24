package webserver

import (
	"github.com/chumnend/pook/app/pook-api/models"
	"github.com/chumnend/pook/app/pook-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

type Server struct {
	Name   string
	DB     *gorm.DB
	Router *gin.Engine
	Config *Config
}

var Pook *Server

func New() (*Server, error) {
	// load config
	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}

	// connect database
	db, err := models.Connect(cfg.DB)

	// setup router
	router := routes.MakeRouter()

	Pook = &Server{
		Name:   "Pook",
		DB:     db,
		Router: router,
		Config: cfg,
	}

	return Pook, nil
}

func (s *Server) Start() {
	s.Router.Run(":" + s.Config.Port)
}
