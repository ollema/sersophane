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

	_ "github.com/lib/pq"
)

type config struct {
	port int
	db   struct {
		dsn string
	}
}

type application struct {
	log           *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("SERSOPHANE_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	addr := fmt.Sprintf(":%d", cfg.port)

	log := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		log:           log,
		templateCache: templateCache,
	}

	log.Printf("Staring server on http://localhost%s", addr)
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
