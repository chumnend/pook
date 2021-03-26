package user

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Attach setups User API
func Attach(r *mux.Router, db *gorm.DB) {
	db.AutoMigrate(&User{})

	userRepo := NewUserRepository(db)
	userUC := NewUserUsecase(userRepo)

	AddUserHandler(r, userUC)
}
