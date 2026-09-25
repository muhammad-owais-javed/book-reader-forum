package handlers

import (
	"context"
	"net/http"
	"time"
)

func (app *Application) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	cookie, err := r.Cookie("session_id")
	if err == nil {
		app.SessionService.DeleteSession(ctx, cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
