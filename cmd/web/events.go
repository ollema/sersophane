package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ollema/sersophane/pkg/forms"
	"github.com/ollema/sersophane/pkg/models"
)

func (app *application) eventCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil || id < 1 {
			app.clientError(w, http.StatusNotFound)
			return
		}

		e, err := app.events.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.clientError(w, http.StatusNotFound)
			} else {
				app.serverError(w, err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyEvent, e)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) eventsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filters, err := models.NewFilters(r.URL.Query(), models.EventSortableColumns, "event_start")
		if err != nil {
			app.clientError(w, http.StatusNotFound)
			return
		}
		events, metadata, err := app.events.GetPage(filters)
		if err != nil {
			app.serverError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyEvents, events)
		ctx = context.WithValue(ctx, contextKeyMetadata, metadata)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) listEvents(w http.ResponseWriter, r *http.Request) {
	events := r.Context().Value(contextKeyEvents).([]*models.Event)
	metadata := r.Context().Value(contextKeyMetadata).(*models.Metadata)
	app.render(w, r, "event.list.page.html", &templateData{Events: events, Metadata: metadata})
}

func (app *application) createEventForm(w http.ResponseWriter, r *http.Request) {
	artists, err := app.artists.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	venues, err := app.venues.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "event.create.page.html", &templateData{Form: forms.New(nil), Artists: artists, Venues: venues})
}

func (app *application) createEvent(w http.ResponseWriter, r *http.Request) {
	artists, err := app.artists.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	venues, err := app.venues.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)

	// break out into form parse/valid func

	var artistIds []int
	var artistIdsStr []string
	for _, artistIdStr := range form.Values["artist"] {
		if artistIdStr != "" {
			artistId, err := strconv.Atoi(artistIdStr)
			if err != nil {
				app.serverError(w, err)
				return
			}

			// only add existing artists, else error out
			artistExists := false
			for _, artist := range artists {
				if artistId == artist.ID {
					artistExists = true
				}
			}
			if !artistExists {
				app.serverError(w, fmt.Errorf("artist with id=%d does not exist", artistId))
				return
			}

			// don't add an artist twice, else skip
			for _, existingArtistId := range artistIds {
				if artistId == existingArtistId {
					break
				}
			}

			artistIds = append(artistIds, artistId)
			artistIdsStr = append(artistIdsStr, artistIdStr)
		}
	}
	form.Values["artist"] = artistIdsStr

	form.Required("artist", "venue", "name", "type", "date", "duration")
	form.MaxLength("name", 100)
	form.ValidDate("date")
	form.ValidNumber("duration", 1, 14)

	if !form.Valid() {
		app.render(w, r, "event.create.page.html", &templateData{Form: form, Artists: artists, Venues: venues})
		return
	}

	venueId, err := strconv.Atoi(form.Get("venue"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	venueExists := false
	for _, venue := range venues {
		if venueId == venue.ID {
			venueExists = true
		}
	}
	if !venueExists {
		app.serverError(w, err)
		return
	}

	var eventType models.EventType
	switch form.Get("type") {
	case string(models.Concert):
		eventType = models.Concert
	case string(models.Festival):
		eventType = models.Festival
	case string(models.Film):
		eventType = models.Film
	default:
		app.serverError(w, fmt.Errorf("%s could not be matched with an EventType", form.Get("type")))
		return
	}

	start, err := time.Parse("2006-01-02", form.Get("date"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	start = start.Add(12 * time.Hour)

	duration, err := strconv.Atoi(form.Get("duration"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	end := start.Add(time.Duration(duration*24) * time.Hour)

	err = app.events.Insert(form.Get("name"), eventType, start, end, artistIds, venueId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", form.Get("name")+" added.")
	http.Redirect(w, r, "/events", http.StatusSeeOther)
}

func (app *application) getEvent(w http.ResponseWriter, r *http.Request) {
	e := r.Context().Value(contextKeyEvent).(*models.Event)
	app.render(w, r, "event.show.page.html", &templateData{Event: e})
}

func (app *application) updateEvent(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deleteEvent(w http.ResponseWriter, r *http.Request) {
}
