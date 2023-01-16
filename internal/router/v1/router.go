package v1

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/controller"
	"github.com/chumnend/pook/internal/middleware"
	"github.com/chumnend/pook/internal/repository"
	"github.com/chumnend/pook/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AttachRouter creates a new gin router
func AttachRouter(h *gin.Engine, db *gorm.DB) {
	v1 := h.Group("/v1")
	v1.Use(middleware.Cors())

	// health check
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// user endpoint
	userRepo := repository.NewUserPostgresRepository(db)
	if err := userRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	userSrv := service.NewUserService(userRepo)
	userCtl := controller.NewUserController(userSrv)

	v1.POST("/register", userCtl.Register)
	v1.POST("/login", userCtl.Login)

	// book endpoint
	bookRepo := repository.NewBookPostgresRepository(db)
	if err := bookRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	bookSrv := service.NewBookService(bookRepo)
	bookCtl := controller.NewBookController(bookSrv)

	v1.GET("/books", bookCtl.ListBooks)
	v1.POST("/books", bookCtl.CreateBook)
	v1.GET("/books/{id:[0-9]+}", bookCtl.GetBook)
	v1.PUT("/books/{id:[0-9]+}", bookCtl.UpdateBook)
	v1.DELETE("/books/{id:[0-9]+}", bookCtl.DeleteBook)

	// page endpoint
	pageRepo := repository.NewPagePostgresRepository(db)
	if err := pageRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	pageSrv := service.NewPageService(pageRepo)
	pageCtl := controller.NewPageController(pageSrv)

	v1.GET("/pages", pageCtl.ListPages)
	v1.POST("/pages", pageCtl.CreatePage)
	v1.GET("/pages/{id:[0-9]+}", pageCtl.GetPage)
	v1.PUT("/pages/{id:[0-9]+}", pageCtl.UpdatePage)
	v1.DELETE("/pages/{id:[0-9]+}", pageCtl.DeletePage)
}
