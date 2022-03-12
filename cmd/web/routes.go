package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ollema/sersophane/ui"
)

func (app *application) routes() http.Handler {
	logFormatter := &middleware.DefaultLogFormatter{Logger: app.logger}
	middleware.DefaultLogger = middleware.RequestLogger(logFormatter)

	r := chi.NewRouter()

	// middleware used for all routes
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(secureHeaders)

	r.Route("/", func(r chi.Router) {
		// middleware used for dynamic routes
		r.Use(app.session.Enable)
		r.Use(csrf)
		r.Use(app.authenticate)

		r.Get("/", app.home)
		r.Get("/about", app.about)

		r.Route("/user", func(r chi.Router) {
			r.With(app.requireAuthentication).Get("/", app.profile)
			r.Get("/signup", app.signupUserForm)
			r.Post("/signup", app.signupUser)
			r.Get("/login", app.loginUserForm)
			r.Post("/login", app.loginUser)
			r.With(app.requireAuthentication).Post("/logout", app.logoutUser)
		})

		r.Route("/events", func(r chi.Router) {
			r.Get("/", app.listEvents)
			r.With(app.requireAuthentication).Get("/create", app.createEventForm)
			r.With(app.requireAuthentication).Post("/create", app.createEvent)

			r.Route("/{id}", func(r chi.Router) {
				r.Use(app.eventCtx)
				r.Get("/", app.getEvent)
				r.With(app.requireAuthentication).Put("/", app.updateEvent)
				r.With(app.requireAuthentication).Delete("/", app.deleteEvent)
			})
		})

		r.Route("/artists", func(r chi.Router) {
			// get all artists as ctx
			r.Group(func(r chi.Router) {
				r.Use(app.artistsCtx)
				r.Get("/", app.listArtists)
				r.With(app.requireAuthentication).Post("/create", app.createArtist)
			})

			// get a single artist as ctx
			r.Route("/{id}", func(r chi.Router) {
				r.Use(app.artistCtx)
				r.Get("/", app.getArtist)
				r.With(app.requireAuthentication).Put("/", app.updateArtist)
				r.With(app.requireAuthentication).Delete("/", app.deleteArtist)
			})
		})

		r.Route("/venues", func(r chi.Router) {
			// get all venues as ctx
			r.Group(func(r chi.Router) {
				r.Use(app.venuesCtx)
				r.Get("/", app.listVenues)
				r.With(app.requireAuthentication).Post("/create", app.createVenue)
			})

			// get a single venue as ctx
			r.Route("/{id}", func(r chi.Router) {
				r.Use(app.venueCtx)
				r.Get("/", app.getVenue)
				r.With(app.requireAuthentication).Put("/", app.updateVenue)
				r.With(app.requireAuthentication).Delete("/", app.deleteVenue)
			})
		})
	})

	fileServer := http.FileServer(http.FS(ui.Files))
	r.Handle("/static/*", fileServer)

	return r
}
