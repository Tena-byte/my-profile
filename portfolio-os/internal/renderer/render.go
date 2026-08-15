package renderer

import (
	"html/template"
	"io/fs"
	"net/http"
)

type Renderer struct {
	templates *template.Template
}

func New(files fs.FS) (*Renderer, error) {
	templates, err := template.ParseFS(
		files,
		"templates/layouts/*.html",
		"templates/pages/*.html",
		"templates/partials/*.html",
	)
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
