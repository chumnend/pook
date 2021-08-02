package page

import (
	"log"

	"github.com/chumnend/pook/internal/domain"
	"github.com/chumnend/pook/internal/routes/page/controller"
	"github.com/chumnend/pook/internal/routes/page/repository"
	"github.com/chumnend/pook/internal/routes/page/service"
	"github.com/jinzhu/gorm"
)

// MakeController builds the page controller givin a db instance
func MakeController(db *gorm.DB) domain.PageController {
	pageRepo := repository.NewPostgresRepository(db)
	if err := pageRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	pageSrv := service.NewService(pageRepo)
	pageCtl := controller.NewController(pageSrv)

	return pageCtl
}
