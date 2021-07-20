package page

import (
	"log"

	"github.com/chumnend/pook/internal/api/page/controller"
	"github.com/chumnend/pook/internal/api/page/repository"
	"github.com/chumnend/pook/internal/api/page/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Attach setups up routes for a passed in Router struct
func Attach(router *mux.Router, db *gorm.DB) {
	pageRepo := repository.NewPostgresRepository(db)
	if err := pageRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	pageSrv := service.NewService(pageRepo)
	pageCtl := controller.NewController(pageSrv)

	router.HandleFunc("/pages", pageCtl.ListPages).Methods("GET")
	router.HandleFunc("/pages", pageCtl.CreatePage).Methods("POST", "OPTIONS")
	router.HandleFunc("/page/{id:[0-9]+}", pageCtl.GetPage).Methods("GET")
	router.HandleFunc("/page/{id:[0-9]+}", pageCtl.UpdatePage).Methods("PUT")
	router.HandleFunc("/page/{id:[0-9]+}", pageCtl.DeletePage).Methods("DELETE")
}
