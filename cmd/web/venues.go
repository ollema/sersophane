package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ollema/sersophane/pkg/forms"
	"github.com/ollema/sersophane/pkg/models"
)

func (app *application) venueCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil || id < 1 {
			app.clientError(w, http.StatusNotFound)
			return
		}

		venue, err := app.venues.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.clientError(w, http.StatusNotFound)
			} else {
				app.serverError(w, err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyVenue, venue)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) venuesCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filters, err := models.NewFilters(r.URL.Query(), models.VenueSortableColumns, "venue_name")
		if err != nil {
			app.clientError(w, http.StatusNotFound)
			return
		}
		venues, metadata, err := app.venues.GetPage(filters)
		if err != nil {
			app.serverError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyVenues, venues)
		ctx = context.WithValue(ctx, contextKeyMetadata, metadata)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) listVenues(w http.ResponseWriter, r *http.Request) {
	venues := r.Context().Value(contextKeyVenues).([]*models.Venue)
	metadata := r.Context().Value(contextKeyMetadata).(*models.Metadata)
	app.render(w, r, "venue.list.page.html", &templateData{Form: &forms.Form{}, Venues: venues, Metadata: metadata})
}

func (app *application) createVenue(w http.ResponseWriter, r *http.Request) {
	venues := r.Context().Value(contextKeyVenues).([]*models.Venue)
	metadata := r.Context().Value(contextKeyMetadata).(*models.Metadata)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name", "city")
	form.MaxLength("name", 100)
	form.MaxLength("city", 100)

	if !form.Valid() {
		app.render(w, r, "venue.list.page.html", &templateData{Form: form, Venues: venues, Metadata: metadata})
		return
	}

	err = app.venues.Insert(form.Get("name"), form.Get("city"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateName) {
			form.Errors.Add("name", "Venue already exists")
			app.render(w, r, "venue.list.page.html", &templateData{Form: form, Venues: venues, Metadata: metadata})
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r, "flash", form.Get("name")+" added.")
	http.Redirect(w, r, "/venues", http.StatusSeeOther)
}

func (app *application) getVenue(w http.ResponseWriter, r *http.Request) {
	venue := r.Context().Value(contextKeyVenue).(*models.Venue)
	filters, err := models.NewFilters(r.URL.Query(), models.EventSortableColumns, "event_start")
	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}
	events, metadata, err := app.events.GetPageForVenue(venue.ID, filters)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "venue.show.page.html", &templateData{Events: events, Metadata: metadata, Venue: venue})
}

func (app *application) updateVenue(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deleteVenue(w http.ResponseWriter, r *http.Request) {
}
