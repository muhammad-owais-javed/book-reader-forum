package handlers

import (
	"context"
	"errors"
	"forum/cmd/web/apperrors"
	"forum/internal/constants"
	"forum/internal/services"
	"forum/internal/uuid"
	"net/http"
	"net/mail"
	"time"
	"unicode/utf8"
)

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	tx, err := app.DB.BeginTx(ctx, nil)
	if err != nil {
		apperrors.ServerError(w, err)
	}
	defer tx.Rollback()

	// ------ input syntax validations --------
	email, err := mail.ParseAddress(r.PostFormValue("email"))
	if err != nil {
		apperrors.ClientError(w, "invalid email")
	}
	password := r.PostFormValue("password")
	if utf8.RuneCountInString(password) < constants.MinPasswordLength {
		apperrors.ClientError(w, "too short password")
		return
	}

	// ------ login --------
	userID, err := services.Authenticate(ctx, tx, email.Address, password)
	if err != nil {
		if errors.Is(err, services.ErrEmailDoesntExist) {
			apperrors.ClientError(w, "email doesnt exist")
			return
		} else if errors.Is(err, services.ErrWrongPassword) {
			apperrors.ClientError(w, "wrong password")
			return
		}
		apperrors.ServerError(w, err)
		return
	}

	// ------ UUID creation --------
	UUID, err := uuid.NewUUID()
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}

	// ------ Session creation --------
	_, _ = userID, UUID
	// sessions table needs to be created
	// http.SetCookie with UUID as the value

}
