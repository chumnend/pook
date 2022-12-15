package webserver

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/chumnend/pook/app/pook-api/routes/album"
	"github.com/chumnend/pook/app/pook-api/routes/ping"
)

type Server struct {
	Name   string
	Router *gin.Engine
	Port   string
	Secret string
}

func New() (*Server, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
		log.Println(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing env: PORT")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("missing env: SECRET_KEY")
	}

	// setup router
	router := gin.Default()

	router.GET("/ping", ping.Pong)

	router.GET("/albums", album.GetAll)
	router.GET("/albums/:id", album.GetByID)
	router.POST("/albums", album.Post)

	s := &Server{
		Name:   "Pook",
		Router: router,
		Port:   port,
		Secret: secret,
	}

	return s, err
}

func (s *Server) Start() {
	s.Router.Run(":" + s.Port)
}
