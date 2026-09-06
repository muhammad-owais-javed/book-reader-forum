package handlers

import (
	"forum/cmd/web/apperrors"
	"html/template"
	"net/http"
)

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("ui/html/register.html")
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}

	tmpl.Execute(w, nil)
}
