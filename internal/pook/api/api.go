package api

import (
	"log"

	"github.com/chumnend/pook/internal/pook/api/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// AttachRouter ...
func AttachRouter(router *mux.Router, db *gorm.DB) {
	// initialize repositories
	userRepo := user.NewPostgresRepository(db)
	if err := userRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	// initialize services
	userSrv := user.NewService(userRepo)

	// initialize controllers
	userCtl := user.NewController(userSrv)

	// setup api subrouter
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/register", userCtl.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/login", userCtl.Login).Methods("POST", "OPTIONS")
}
