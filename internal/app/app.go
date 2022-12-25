package app

import (
	"log"

	"github.com/chumnend/pook/config"
	v1 "github.com/chumnend/pook/internal/router/v1"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm Postgres Driver
)

type App struct {
	Cfg    *config.Config
	Db     *gorm.DB
	Router *gin.Engine
}

func New(cfg *config.Config) *App {
	// connect to database
	pg, err := gorm.Open("postgres", cfg.DB)
	if err != nil {
		log.Fatalf("postgres connect error: %s", err)
	}
	defer pg.Close()

	// setup routes
	router := gin.Default()
	v1.AttachRouter(router, pg)

	return &App{
		Cfg:    cfg,
		Db:     pg,
		Router: router,
	}
}

func (app *App) Run() {
	app.Router.Run(":" + app.Cfg.Port)
}
