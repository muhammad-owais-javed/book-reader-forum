package pages

import (
	"forum/cmd/web/apperrors"
	"html/template"
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("ui/html/login.html")
	if err != nil {
		apperrors.ServerError(w, err)
		return
	}

	tmpl.Execute(w, nil)
}
