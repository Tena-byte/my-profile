package renderer

import (
	"html/template"
	"net/http"
)

type Renderer struct {
	templates *template.Template
}

func New() (*Renderer, error) {
	templates, err := template.ParseGlob("templates/layouts/*.html")
	if err != nil {
		return nil, err
	}

	templates, err = templates.ParseGlob("templates/pages/*.html")
	if err != nil {
		return nil, err
	}

	return &Renderer{
		templates: templates,
	}, nil
}

func (r *Renderer) Render(
	w http.ResponseWriter,
	name string,
	data any,
) error {
	return r.templates.ExecuteTemplate(w, name, data)
}