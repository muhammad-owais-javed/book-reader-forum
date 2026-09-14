package handlers

import (
	"net/http"
)

func (app *Application) Routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /register", RegisterPageHandler)
	mux.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		withTimeout(w, r, app.RegistrationHandler)
	})

	mux.HandleFunc("GET /login", LoginPageHandler)
	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		withTimeout(w, r, app.LoginHandler)
	})

	mux.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		withAuthentication(w, r, app.HomePageHandler, app)
	})

	mux.HandleFunc("/forum", func(w http.ResponseWriter, r *http.Request ) {
		withAuthentication(w, r, app.Forum.HelloWorld, app)
	})

	return mux
}
