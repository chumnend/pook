package user

import (
	"log"

	"github.com/chumnend/pook/internal/domain"

	"github.com/chumnend/pook/internal/router/user/controller"
	"github.com/chumnend/pook/internal/router/user/repository"
	"github.com/chumnend/pook/internal/router/user/service"
	"github.com/jinzhu/gorm"
)

// MakeController builds the user controller givin a db instance
func MakeController(db *gorm.DB) domain.UserController {
	userRepo := repository.NewPostgresRepository(db)
	if err := userRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	userSrv := service.NewService(userRepo)
	userCtl := controller.NewController(userSrv)

	return userCtl
}
