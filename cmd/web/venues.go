package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ollema/sersophane/pkg/forms"
	"github.com/ollema/sersophane/pkg/models"
	"github.com/ollema/sersophane/pkg/models/postgres"
)

func (app *application) venueCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil || id < 1 {
			app.clientError(w, http.StatusNotFound)
			return
		}

		e, err := app.venues.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.clientError(w, http.StatusNotFound)
			} else {
				app.serverError(w, err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyVenue, e)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) listVenues(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	filters := postgres.NewFilters(page, "name", "ASC")

	venues, err := app.venues.GetAll(filters)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "venue.list.page.html", &templateData{Form: &forms.Form{}, Venues: venues})
}

func (app *application) createVenue(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name")
	form.MaxLength("name", 100)

	if !form.Valid() {
		app.render(w, r, "venue.list.page.html", &templateData{Form: form})
		return
	}

	err = app.venues.Insert(form.Get("name"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateName) {
			form.Errors.Add("name", "Venue already exists")
			app.render(w, r, "venue.list.page.html", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r, "flash", form.Get("name")+" added.")
	http.Redirect(w, r, "/venues", http.StatusSeeOther)
}

func (app *application) getVenue(w http.ResponseWriter, r *http.Request) {
	v := r.Context().Value(contextKeyVenue).(*models.Venue)
	app.render(w, r, "venue.show.page.html", &templateData{Venue: v})
}

func (app *application) updateVenue(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deleteVenue(w http.ResponseWriter, r *http.Request) {
}
