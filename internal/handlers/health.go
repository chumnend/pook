package handlers

import (
	"log"
	"net/http"

	"github.com/chumnend/pook/internal/utils"
)

func CheckHealth(w http.ResponseWriter, req *http.Request) {
	log.Println("Request made to" + req.URL.Path)

	response := map[string]string{
		"status": "ok",
	}
	utils.SendJSON(w, response, http.StatusOK)
}
