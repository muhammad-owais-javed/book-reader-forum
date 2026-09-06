package apperrors

import (
	"fmt"
	"forum/internal/constants"
	"log"
	"net/http"
)

// global error handling for bad requests of different kind
func ClientError(w http.ResponseWriter, errorName string) {

	errorMessage, statusCode := WriteClientError(errorName)
	http.Error(w, errorMessage, statusCode)

}

// writes the error message and status code depending which error it receives
func WriteClientError(errorName string) (string, int) {

	switch errorName {
	case "bad request":
		return http.StatusText(http.StatusBadRequest), http.StatusBadRequest

	case "invalid email":
		return "Error: Invalid email provided.", http.StatusBadRequest

	case "too short password":
		return fmt.Sprintf("Error: The provided password is too short. Minimum length %d characters.", constants.MinPasswordLength), http.StatusBadRequest

	case "too short username":
		return fmt.Sprintf("Error: The provided username is too short. Minimum length %d characters.", constants.MinUsernameLength), http.StatusBadRequest

	case "email exists":
		return "Error: The provided email address has already a user.", http.StatusConflict

	case "username exists":
		return "Error: The provided username is already in use.", http.StatusConflict

	case "unauthorized":
		return http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized

	default:
		log.Println("We should not have gotten here.")
		return "", 0
	}

}

// global error handling for server side issues
func ServerError(w http.ResponseWriter, err error) {

	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
