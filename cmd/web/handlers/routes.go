package handlers

import (
	"forum/cmd/web/handlers/pages"
	"net/http"
)

func (app *Application) Routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /register", pages.RegisterPageHandler)
	mux.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		withTimeout(w, r, app.RegistrationHandler)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		withTimeout(w, r, app.LoginHandler)
	})

	return mux
}
