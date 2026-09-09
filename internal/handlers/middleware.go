package handlers

import (
	"context"
	"forum/internal/errors"
	"net/http"
	"time"
)

// wraps a timeout condition around the handler
func withTimeout(w http.ResponseWriter, r *http.Request, handler http.HandlerFunc) {

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	handler(w, r.WithContext(ctx)) // calling the handler that was passed, but with context
}

var Authenticated bool

// requires user to authenticate
func withAuthentication(w http.ResponseWriter, r *http.Request, handler http.HandlerFunc, app *Application) {

	//ctx and tx necessary to call service level functions, so they are defined/added here to satisfy syntax
	ctx := r.Context()
	tx, err := app.DB.BeginTx(ctx, nil)
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}
	defer tx.Rollback()

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther) // 303, "Go check this out!"
		return
	}
	sessionID := cookie.Value

	validSession, err := app.SessionService.ValidateSession(ctx, sessionID)
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}
	if !validSession {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tx.Commit()

	handler(w, r)
}
