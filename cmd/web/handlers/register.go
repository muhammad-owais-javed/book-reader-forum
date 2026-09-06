package handlers

import (
	"context"
	"errors"
	"forum/internal/constants"
	"forum/internal/services"
	"net/http"
	"net/mail"
	"time"
	"unicode/utf8"
)

func (app *Application) RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	username := r.PostFormValue("username")
	if utf8.RuneCountInString(username) < constants.MinUsernameLength {
		clientError(w, "too short username")
		return
	}

	email, err := mail.ParseAddress(r.PostFormValue("email"))
	if err != nil {
		clientError(w, "invalid email")
		return
	}

	password := r.PostFormValue("password")
	if utf8.RuneCountInString(password) < constants.MinPasswordLength {
		clientError(w, "too short password")
		return
	}

	err = app.Registration.Register(ctx, username, email.Address, password)
	if err != nil {
		if errors.Is(err, services.ErrEmailExists) {
			clientError(w, "email exists")
			return
		} else if errors.Is(err, services.ErrUsernameExists) {
			clientError(w, "username exists")
			return
		}
		serverError(w, err)
		return
	}

}
