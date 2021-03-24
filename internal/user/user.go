package user

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/chumnend/pook/internal/user/domain"
	"github.com/chumnend/pook/internal/user/entity"
	"github.com/chumnend/pook/internal/user/handler"
	"github.com/chumnend/pook/internal/user/repository"
)

// Attach setups User API
func Attach(r *mux.Router, db *gorm.DB) {
	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserRepository(db)
	userEntity := entity.NewUserEntity(userRepo)

	handler.NewUserHandler(r, userEntity)
}
