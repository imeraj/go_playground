package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	logging "github.com/imeraj/go_playground/gowiki/middleware"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("templates/view.html", "templates/edit.html"))
var validTitle = regexp.MustCompile("^([a-zA-Z0-9]+)$")

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := validTitle.FindStringSubmatch(vars["title"])
	if t == nil {
		http.NotFound(w, r)
		return
	}
	title := t[0]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := validTitle.FindStringSubmatch(vars["title"])
	if t == nil {
		http.NotFound(w, r)
		return
	}
	title := t[0]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := validTitle.FindStringSubmatch(vars["title"])
	if t == nil {
		http.NotFound(w, r)
		return
	}
	title := t[0]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/view/{title}", logging.Chain(viewHandler, logging.Logging())).Methods("GET")
	router.HandleFunc("/edit/{title}", logging.Chain(editHandler, logging.Logging())).Methods("GET")
	router.HandleFunc("/save/{title}", logging.Chain(saveHandler, logging.Logging())).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
