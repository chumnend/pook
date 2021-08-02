package book

import (
	"log"

	"github.com/chumnend/pook/internal/domain"
	"github.com/chumnend/pook/internal/routes/book/controller"
	"github.com/chumnend/pook/internal/routes/book/repository"
	"github.com/chumnend/pook/internal/routes/book/service"
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
