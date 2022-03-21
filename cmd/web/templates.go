package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
	"getArtistIdFromForm":  getArtistIdFromForm,
	"N":                    N,
	"shortDate":            shortDate,
	"toInt":                toInt,
	"toLower":              strings.ToLower,
	"title":                title,
	"metadataToPrevPage":   metadataToPrevPage,
	"metadataToPages":      metadataToPages,
	"metadataToNextPage":   metadataToNextPage,
	"metadataToStartRange": metadataToStartRange,
	"metadataToEndRange":   metadataToEndRange,
}

func N(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			stream <- i
		}
		close(stream)
	}()
	return
}

func getArtistIdFromForm(form forms.Form, index int) int {
	if index < len(form.Values["artist"]) {
		artistId, err := strconv.Atoi(form.Values["artist"][index])
		if err != nil {
			return -1
		}
		return artistId
	}
	return -1
}

func shortDate(d time.Time) string {
	return fmt.Sprintf("%d/%d %d", d.Day(), d.Month(), d.Year())
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}

func title(input string) string {
	words := strings.Fields(input)
	smallwords := " a an of on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func metadataToPrevPage(metadata *models.Metadata) int {
	return int(math.Max(float64(metadata.CurrentPage-1), float64(metadata.FirstPage)))
}

func metadataToPages(metadata *models.Metadata) []int {
	var maxEndDiff int
	switch metadata.CurrentPage {
	case 1:
		maxEndDiff = 4
	case 2:
		maxEndDiff = 3
	default:
		maxEndDiff = 2
	}

	var maxStartDiff int
	switch metadata.LastPage - metadata.CurrentPage {
	case 0:
		maxStartDiff = 4
	case 1:
		maxStartDiff = 3
	default:
		maxStartDiff = 2
	}

	start := int(math.Max(float64(metadata.CurrentPage-maxStartDiff), float64(metadata.FirstPage)))
	end := int(math.Min(float64(metadata.CurrentPage+maxEndDiff), float64(metadata.LastPage)))

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
