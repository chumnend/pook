package api

import (
	"github.com/chumnend/pook/internal/api/book"
	"github.com/chumnend/pook/internal/api/page"
	"github.com/chumnend/pook/internal/api/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// SetupRoutes is used to setup the api for the pook app
func SetupRoutes(router *mux.Router, db *gorm.DB) {
	api := router.PathPrefix("/api/v1").Subrouter()
	user.Attach(api, db)
	book.Attach(api, db)
	page.Attach(api, db)
}
