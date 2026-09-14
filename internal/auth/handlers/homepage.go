package handlers

import (
	"forum/internal/errors"
	"html/template"
	"net/http"
)

func (app *Application) HomePageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("ui/html/home.html")
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}

	tmpl.Execute(w, nil)
}
