package api

import (
	"log"

	"github.com/chumnend/pook/internal/pook/api/book"
	"github.com/chumnend/pook/internal/pook/api/page"
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

	bookRepo := book.NewPostgresRepository(db)
	if err := bookRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	pageRepo := page.NewPostgresRepository(db)
	if err := pageRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	// initialize services
	userSrv := user.NewService(userRepo)
	bookSrv := book.NewService(bookRepo)
	pageSrv := page.NewService(pageRepo)

	// initialize controllers
	userCtl := user.NewController(userSrv)
	bookCtl := book.NewController(bookSrv)
	pageCtl := page.NewController(pageSrv)

	// setup api subrouter
	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/register", userCtl.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/login", userCtl.Login).Methods("POST", "OPTIONS")

	api.HandleFunc("/books", bookCtl.ListBooks).Methods("GET")
	api.HandleFunc("/books", bookCtl.CreateBook).Methods("POST", "OPTIONS")
	api.HandleFunc("/book/{id:[0-9]+}", bookCtl.GetBook).Methods("GET")
	api.HandleFunc("/book/{id:[0-9]+}", bookCtl.UpdateBook).Methods("PUT")
	api.HandleFunc("/book/{id:[0-9]+}", bookCtl.DeleteBook).Methods("DELETE")

	api.HandleFunc("/pages", pageCtl.ListPages).Methods("GET")
	api.HandleFunc("/pages", pageCtl.CreatePage).Methods("POST", "OPTIONS")
	api.HandleFunc("/page/{id:[0-9]+}", pageCtl.GetPage).Methods("GET")
	api.HandleFunc("/page/{id:[0-9]+}", pageCtl.UpdatePage).Methods("PUT")
	api.HandleFunc("/page/{id:[0-9]+}", pageCtl.DeletePage).Methods("DELETE")
}
