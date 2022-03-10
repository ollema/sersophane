package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.html", &templateData{ActiveNavItem: HomeNavItem})
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.page.html", &templateData{ActiveNavItem: AboutNavItem})
}
