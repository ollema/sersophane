package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ollema/sersophane/pkg/models"
	"github.com/ollema/sersophane/pkg/models/postgres"
)

type application struct {
	artists interface {
		Insert(string) error
		Get(int) (*models.Artist, error)
		GetPage(*models.Filters) ([]*models.Artist, *models.Metadata, error)
		GetAll() ([]*models.Artist, error)
	}
	events interface {
		Insert(string, models.EventType, time.Time, time.Time, []int, int) error
		Get(int) (*models.Event, error)
		GetPage(*models.Filters) ([]*models.Event, *models.Metadata, error)
		GetPageForArtist(int, *models.Filters) ([]*models.Event, *models.Metadata, error)
		GetPageForVenue(int, *models.Filters) ([]*models.Event, *models.Metadata, error)
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
		Insert(string, string) error
		Get(int) (*models.Venue, error)
		GetPage(*models.Filters) ([]*models.Venue, *models.Metadata, error)
		GetAll() ([]*models.Venue, error)
	}
}

func main() {
	var port int
	var dsn string
	var secret string
	flag.IntVar(&port, "port", 4000, "port")
	flag.StringVar(&dsn, "db-dsn", os.Getenv("SERSOPHANE_DB_DSN"), "postgres dsn")
	flag.StringVar(&secret, "secret", os.Getenv("SERSOPHANE_SECRET"), "secret (for encrypting session)")
	flag.Parse()

	addr := fmt.Sprintf(":%d", port)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db := openDB(dsn)
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	session := sessions.New([]byte(secret))
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

func openDB(dsn string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}
