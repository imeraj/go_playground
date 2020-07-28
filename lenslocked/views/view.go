package views

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/gorilla/csrf"
	"github.com/imeraj/go_playground/lenslocked/context"
)

type View struct {
	Template *template.Template
	Layout   string
}

const (
	LayoutDir   = "views/layouts/"
	TemplateDir = "views/"
	TemplateExt = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)

	files = append(files, layoutFiles()...) // unpack slice to variadic parameters
	t, err := template.New("").Funcs(template.FuncMap{
		"pathEscape": func(s string) string {
			return url.PathEscape(s)
		},
		"csrfField": func() (template.HTML, error) {
			return "", errors.New("csrfField is not implemented")
		},
	}).ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{Template: t, Layout: layout}
}

func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) error {
	w.Header().Add("Content-Type", "text/html")
	var vd Data
	switch d := data.(type) {
	case Data:
		vd = d
	default:
		vd = Data{
			Yield: data,
		}
	}
	vd.User = context.User(r.Context())

	var buf bytes.Buffer
	csrfField := csrf.TemplateField(r)
	tpl := v.Template.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrfField
		},
	})
	err := tpl.ExecuteTemplate(&buf, v.Layout, vd)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return nil
	}
	io.Copy(w, &buf)
	return nil
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, r, nil); err != nil {
		panic(err)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
