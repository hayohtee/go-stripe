package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]any
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated bool
	API             string
	CSSVersion      string
}

var functions = template.FuncMap{}

//go:embed templates
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateName := fmt.Sprintf("templates/%s_page.gohtml", page)

	_, ok := app.templateCache[templateName]
	if app.cfg.env == "production" && ok {
		t = app.templateCache[templateName]
	} else {
		t, err = app.parseTemplate(partials, page, templateName)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) parseTemplate(partials []string, page, templateName string) (*template.Template, error) {
	var t *template.Template
	var err error

	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s_partial.gohtml", x)
		}

		t, err = template.New(fmt.Sprintf("%s_page.gohtml", page)).Funcs(functions).ParseFS(
			templateFS,
			"templates/base_layout.gohtml",
			strings.Join(partials, ","),
			templateName,
		)
	} else {
		t, err = template.New(fmt.Sprintf("%s_page.gohtml", page)).Funcs(functions).ParseFS(
			templateFS,
			"templates/base_layout.gohtml",
			templateName,
		)
	}

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateName] = t
	return t, nil
}
