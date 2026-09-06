package handlers

import (
	"fmt"
	"net/http"
)

func clientError(w http.ResponseWriter, status int) {

	switch status {
	case 400:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	case 401:
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	default:
		fmt.Println("We should not have gotten here.")
	}

}
