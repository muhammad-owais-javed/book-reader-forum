package handlers

import (
	"context"
	"net/http"
	"time"
)

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	userID, err := app.Auth.Authenticate(ctx, email, password)
	if err != nil {
		// implement global error handling here
	}

	_ = userID // --------
	// activate generation of UUID, which will be sent to the browser
	// sessions table needs to be created
	// http.SetCookie with UUID as the value

}
