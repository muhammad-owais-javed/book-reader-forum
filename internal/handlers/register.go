package handlers

import (
	"context"
	"errors"
	"forum/internal/errors"
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

	tx, err := app.DB.BeginTx(ctx, nil)
	if err != nil {
		apperrors.ServerError(w, err)
	}
	defer tx.Rollback()

	// ------ input syntax validations --------
	username := r.PostFormValue("username")
	if utf8.RuneCountInString(username) < constants.MinUsernameLength {
		apperrors.ClientError(w, "too short username")
		return
	}

	email, err := mail.ParseAddress(r.PostFormValue("email"))
	if err != nil {
		apperrors.ClientError(w, "invalid email")
		return
	}

	password := r.PostFormValue("password")
	if utf8.RuneCountInString(password) < constants.MinPasswordLength {
		apperrors.ClientError(w, "too short password")
		return
	}

	// ------ registration --------
	err = services.Register(ctx, tx, username, email.Address, password)
	if err != nil {
		if errors.Is(err, services.ErrEmailExists) {
			apperrors.ClientError(w, "email exists")
			return
		} else if errors.Is(err, services.ErrUsernameExists) {
			apperrors.ClientError(w, "username exists")
			return
		}
		apperrors.ServerError(w, err)
		return
	}

	tx.Commit()

	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
