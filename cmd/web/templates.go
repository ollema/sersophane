package main

import (
	"html/template"
	"io/fs"
	"math"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/justinas/nosurf"
	"github.com/ollema/sersophane/pkg/forms"
	"github.com/ollema/sersophane/pkg/models"
	"github.com/ollema/sersophane/ui"
)

type templateData struct {
	Artist          *models.Artist
	Artists         []*models.Artist
	CSRFToken       string
	Event           *models.Event
	Events          []*models.Event
	Flash           string
	Form            *forms.Form
	IsAuthenticated bool
	Metadata        *models.Metadata
	NavItems        []NavItem
	ActiveNavItem   NavItem
	User            *models.User
	Venue           *models.Venue
	Venues          []*models.Venue
}

type NavItem struct {
	Name string
	Link string
}

var DefaultNavItems []NavItem = []NavItem{
	HomeNavItem,
	EventsNavItem,
	ArtistsNavItem,
	VenuesNavItem,
	AboutNavItem,
}

var HomeNavItem = NavItem{"Home", "/"}
var EventsNavItem = NavItem{"Events", "/events"}
var ArtistsNavItem = NavItem{"Artists", "/artists"}
var VenuesNavItem = NavItem{"Venues", "/venues"}
var AboutNavItem = NavItem{"About", "/about"}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CSRFToken = nosurf.Token(r)
	td.Flash = app.session.PopString(r, "flash")
	td.IsAuthenticated = app.isAuthenticated(r)
	td.NavItems = DefaultNavItems

	if td.IsAuthenticated {
		user, err := app.users.Get(app.session.GetInt(r, "authenticatedUserID"))
		if err == nil {
			td.User = &models.User{Name: user.Name, Email: user.Email}
		}
	}

	return td
}

var functions = template.FuncMap{
	"toLower":              strings.ToLower,
	"title":                strings.Title,
	"metadataToPrevPage":   metadataToPrevPage,
	"metadataToPages":      metadataToPages,
	"metadataToNextPage":   metadataToNextPage,
	"metadataToStartRange": metadataToStartRange,
	"metadataToEndRange":   metadataToEndRange,
}

func metadataToPrevPage(metadata *models.Metadata) int {
	return int(math.Max(float64(metadata.CurrentPage-1), float64(metadata.FirstPage)))
}

func metadataToPages(metadata *models.Metadata) []int {
	start := int(math.Max(float64(metadata.CurrentPage-2), float64(metadata.FirstPage)))
	end := int(math.Min(float64(metadata.CurrentPage+2), float64(metadata.LastPage)))

	pages := make([]int, end-start+1)
	for i := range pages {
		pages[i] = start + i
	}

	return pages
}

func metadataToNextPage(metadata *models.Metadata) int {
	return int(math.Min(float64(metadata.CurrentPage+1), float64(metadata.LastPage)))
}

func metadataToStartRange(metadata *models.Metadata) int {
	return (metadata.CurrentPage-1)*metadata.PageSize + 1
}

func metadataToEndRange(metadata *models.Metadata) int {
	return int(math.Min(float64(metadata.CurrentPage*metadata.PageSize), float64(metadata.TotalRecords)))
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/*.page.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFS(ui.Files, "html/*.layout.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFS(ui.Files, "html/*.partial.html")
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
