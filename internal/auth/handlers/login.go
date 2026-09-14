package handlers

import (
	"context"
	"errors"
	"forum/internal/errors"
	"forum/internal/constants"
	"forum/internal/auth/services"
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
	userID, err := app.UserService.Authenticate(ctx, email.Address, password)
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

	// ------ session creation --------
	UUID, err := app.SessionService.CreateSession(ctx, userID)
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: UUID,
	})

	tx.Commit()

	//http.Redirect(w, r, "/home", http.StatusSeeOther)
	http.Redirect(w, r, "/forum", http.StatusSeeOther )

}
