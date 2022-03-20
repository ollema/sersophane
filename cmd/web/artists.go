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

func (app *application) artistCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil || id < 1 {
			app.clientError(w, http.StatusNotFound)
			return
		}

		e, err := app.artists.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.clientError(w, http.StatusNotFound)
			} else {
				app.serverError(w, err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyArtist, e)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) artistsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filters, err := models.NewFilters(r.URL.Query(), models.ArtistSortableColumns, "artist_name")
		if err != nil {
			app.clientError(w, http.StatusNotFound)
			return
		}
		artists, metadata, err := app.artists.GetPage(filters)
		if err != nil {
			app.serverError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyArtists, artists)
		ctx = context.WithValue(ctx, contextKeyMetadata, metadata)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) listArtists(w http.ResponseWriter, r *http.Request) {
	artists := r.Context().Value(contextKeyArtists).([]*models.Artist)
	metadata := r.Context().Value(contextKeyMetadata).(*models.Metadata)
	app.render(w, r, "artist.list.page.html", &templateData{Form: &forms.Form{}, Artists: artists, Metadata: metadata})
}

func (app *application) createArtist(w http.ResponseWriter, r *http.Request) {
	artists := r.Context().Value(contextKeyArtists).([]*models.Artist)
	metadata := r.Context().Value(contextKeyMetadata).(*models.Metadata)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name")
	form.MaxLength("name", 100)

	if !form.Valid() {
		app.render(w, r, "artist.list.page.html", &templateData{Form: form, Artists: artists, Metadata: metadata})
		return
	}

	err = app.artists.Insert(form.Get("name"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateName) {
			form.Errors.Add("name", "Artist already exists")
			app.render(w, r, "artist.list.page.html", &templateData{Form: form, Artists: artists, Metadata: metadata})
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r, "flash", form.Get("name")+" added.")
	http.Redirect(w, r, "/artists", http.StatusSeeOther)
}

func (app *application) getArtist(w http.ResponseWriter, r *http.Request) {
	artist := r.Context().Value(contextKeyArtist).(*models.Artist)
	filters, err := models.NewFilters(r.URL.Query(), models.EventSortableColumns, "event_start")
	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}
	events, metadata, err := app.events.GetPageForArtist(artist.ID, filters)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "artist.show.page.html", &templateData{Artist: artist, Events: events, Metadata: metadata})
}

func (app *application) updateArtist(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deleteArtist(w http.ResponseWriter, r *http.Request) {
}
