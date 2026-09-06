package handlers

import (
	"log"
	"net/http"
)

// global error handling for bad requests of different kind
func clientError(w http.ResponseWriter, errorName string) {

	errorMessage, statusCode := writeClientError(errorName)
	http.Error(w, errorMessage, statusCode)

}

// writes the error message and status code depending which error it receives
func writeClientError(errorName string) (string, int) {

	switch errorName {
	case "bad request":
		return http.StatusText(http.StatusBadRequest), http.StatusBadRequest
	case "invalid email":
		return "Error: Invalid email provided.", http.StatusBadRequest
	case "unauthorized":
		return http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized

	default:
		log.Println("We should not have gotten here.")
		return "", 0
	}

}

// global error handling for server side issues
func serverError(w http.ResponseWriter, err error) {

	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
