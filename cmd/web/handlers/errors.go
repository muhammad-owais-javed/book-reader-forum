package handlers

import (
	"log"
	"net/http"
)

// global error handling for bad request of different forms
func clientError(w http.ResponseWriter, status int) {

	switch status {
	case 400:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	case 401:
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	default:
		log.Println("We should not have gotten here.")
	}
}

// global error handling for server side issues
func serverError(w http.ResponseWriter, err error) {

	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
