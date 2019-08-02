package views

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type View struct {
	Template *template.Template
	Layout   string
}

const (
	LayoutDir   = "views/layouts/"
	TemplateExt = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...) // unpack slice to variadic parameters
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{Template: t, Layout: layout}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
