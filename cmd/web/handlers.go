package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.html", &templateData{})
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	return td
}
