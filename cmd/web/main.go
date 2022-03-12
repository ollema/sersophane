package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	_ "github.com/lib/pq"
	"github.com/ollema/sersophane/pkg/models"
	"github.com/ollema/sersophane/pkg/models/postgres"
)

type config struct {
	db struct {
		dsn string
	}
	port   int
	secret string
}

type application struct {
	artists interface {
		Insert(string) error
		Get(int) (*models.Artist, error)
	}
	events interface {
		Insert(string, models.EventType, time.Time, time.Time) error
		Get(int) (*models.Event, error)
	}
	logger        *log.Logger
	session       *sessions.Session
	templateCache map[string]*template.Template
	users         interface {
		Insert(string, string, string) error
		Authenticate(string, string) (int, error)
		Get(int) (*models.User, error)
	}
	venues interface {
		Insert(string) error
		Get(int) (*models.Venue, error)
		GetAll(filters postgres.Filters) ([]*models.Venue, error)
	}
}

type contextKey string

const (
	contextKeyEvent           contextKey = "event"
	contextKeyArtist          contextKey = "artist"
	contextKeyIsAuthenticated contextKey = "isAuthenticated"
	contextKeyVenue           contextKey = "venue"
)

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("SERSOPHANE_DB_DSN"), "postgres dsn")
	flag.StringVar(&cfg.secret, "secret", os.Getenv("SERSOPHANE_SECRET"), "secret (for encrypting session)")
	flag.Parse()

	addr := fmt.Sprintf(":%d", cfg.port)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	session := sessions.New([]byte(cfg.secret))
	session.Lifetime = 12 * time.Hour
	// session.Secure = true  TODO: set true in prod

	app := &application{
		artists:       &postgres.ArtistModel{DB: db},
		events:        &postgres.EventModel{DB: db},
		logger:        logger,
		session:       session,
		templateCache: templateCache,
		users:         &postgres.UserModel{DB: db},
		venues:        &postgres.VenueModel{DB: db},
	}

	log.Printf("Starting server on http://localhost%s", addr)
	err = http.ListenAndServe(addr, app.routes())
	log.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
