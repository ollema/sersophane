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

func (app *application) listEvents(w http.ResponseWriter, r *http.Request) {
}

func (app *application) createEventForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "event.create.page.html", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createEvent(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getEvent(w http.ResponseWriter, r *http.Request) {
	e := r.Context().Value(contextKeyEvent).(*models.Event)
	app.render(w, r, "event.show.page.html", &templateData{Event: e})
}

func (app *application) updateEvent(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deleteEvent(w http.ResponseWriter, r *http.Request) {
}
