package book

import (
	"log"

	"github.com/chumnend/pook/internal/api/book/controller"
	"github.com/chumnend/pook/internal/api/book/repository"
	"github.com/chumnend/pook/internal/api/book/service"
	"github.com/chumnend/pook/internal/domain"
	"github.com/jinzhu/gorm"
)

// MakeController builds the book controller givin a db instance
func MakeController(db *gorm.DB) domain.BookController {
	bookRepo := repository.NewPostgresRepository(db)
	if err := bookRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	bookSrv := service.NewService(bookRepo)
	bookCtl := controller.NewController(bookSrv)

	return bookCtl
}
