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
			r.Get("/create", app.createEventForm)
			r.Post("/create", app.createEvent)

			r.Route("/{id}", func(r chi.Router) {
				r.Use(app.eventCtx)
				r.Get("/", app.getEvent)
				r.Put("/", app.updateEvent)
				r.Delete("/", app.deleteEvent)
			})
		})
	})

	fileServer := http.FileServer(http.FS(ui.Files))
	r.Handle("/static/*", fileServer)

	return r
}
