package book

import (
	"log"

	"github.com/chumnend/pook/internal/book/controller"
	"github.com/chumnend/pook/internal/book/repository"
	"github.com/chumnend/pook/internal/book/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Attach setups up routes for a passed in Router struct
func Attach(router *mux.Router, db *gorm.DB) {
	bookRepo := repository.NewPostgresRepository(db)
	if err := bookRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	bookSrv := service.NewService(bookRepo)
	bookCtl := controller.NewController(bookSrv)

	router.HandleFunc("/books", bookCtl.ListBooks).Methods("GET")
	router.HandleFunc("/books", bookCtl.CreateBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/book/{id:[0-9]+}", bookCtl.GetBook).Methods("GET")
	router.HandleFunc("/book/{id:[0-9]+}", bookCtl.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id:[0-9]+}", bookCtl.DeleteBook).Methods("DELETE")
}
