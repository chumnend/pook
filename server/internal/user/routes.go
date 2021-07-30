package user

import (
	"log"

	"github.com/chumnend/pook/server/internal/user/controller"
	"github.com/chumnend/pook/server/internal/user/repository"
	"github.com/chumnend/pook/server/internal/user/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Attach setups up routes for a passed in Router struct
func Attach(router *mux.Router, db *gorm.DB) {
	userRepo := repository.NewPostgresRepository(db)
	if err := userRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	userSrv := service.NewService(userRepo)
	userCtl := controller.NewController(userSrv)

	router.HandleFunc("/register", userCtl.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", userCtl.Login).Methods("POST", "OPTIONS")
}
