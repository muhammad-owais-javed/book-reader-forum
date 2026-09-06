package handlers

import (
	"context"
	"forum/internal/constants"
	"net/http"
	"net/mail"
	"time"
	"unicode/utf8"
)

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	email, err := mail.ParseAddress(r.PostFormValue("email"))
	if err != nil {
		clientError(w, "invalid email")
	}
	password := r.PostFormValue("password")
	if utf8.RuneCountInString(password) < constants.MinPasswordLength {
		clientError(w, "too short password")
		return
	}

	userID, err := app.Auth.Authenticate(ctx, email.Address, password)
	// use global error handling here

	_ = userID // --------
	// activate generation of UUID, which will be sent to the browser
	// sessions table needs to be created
	// http.SetCookie with UUID as the value

}
