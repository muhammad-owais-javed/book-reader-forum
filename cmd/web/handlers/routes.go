package handlers

import "net/http"

func (app *Application) Routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		withTimeout(w, r, app.LoginHandler)
	})

	return mux
}
