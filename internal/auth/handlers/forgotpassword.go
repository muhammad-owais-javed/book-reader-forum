package handlers

import (
	"context"
	apperrors "forum/internal/errors"
	"html/template"
	"net/http"
	"net/mail"
	"time"
)

func (app *Application) ForgotPasswordPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("ui/html/forgotpassword.html")
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}

	tmpl.Execute(w, nil)
}

func (app *Application) ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	// ------ input syntax validations --------
	email, err := mail.ParseAddress(r.PostFormValue("email"))
	if err != nil {
		apperrors.ClientError(w, "invalid email")
	}

	err = app.ResetTokenService.CreateResetToken(ctx, email.Address)
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}

	tmpl, err := template.ParseFiles("ui/html/resetlinksent.html")
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}

	tmpl.Execute(w, nil)
}
