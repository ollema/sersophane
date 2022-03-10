package main

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/justinas/nosurf"
	"github.com/ollema/sersophane/pkg/forms"
	"github.com/ollema/sersophane/pkg/models"
	"github.com/ollema/sersophane/ui"
)

type templateData struct {
	CSRFToken       string
	Flash           string
	Form            *forms.Form
	IsAuthenticated bool
	NavItems        []NavItem
	ActiveNavItem   NavItem
	User            *models.User
}

type NavItem struct {
	Name string
	Link string
}

var DefaultNavItems []NavItem = []NavItem{
	HomeNavItem,
	AboutNavItem,
}

var HomeNavItem = NavItem{"Home", "/"}
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
	"toLower": strings.ToLower,
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
