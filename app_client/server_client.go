package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/bmizerany/pat"
	"github.com/unrolled/render"
	"github.com/viniciusfeitosa/TalkRubyConf/app_client/handlers"
)

const (
	port string = ":8880"
)

var r *render.Render

func init() {
	r = render.New(render.Options{})
}

func main() {
	superheroHandler := handlers.NewSuperheroHandler(r)

	mux := pat.New()
	mux.Get("/superheros.json", http.HandlerFunc(superheroHandler.JSON))
	mux.Get("/superheros.html", superheroHandler)
	// mux.Get("/", &templateHandler{filename: "index.html"})
	mux.Get("/static/", assets("static"))

	log.Println("Server Superhero start on port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

// templ represents a single template
type templateHandler struct {
	filename string
	templ    *template.Template
}

// func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
// 	t.templ.Execute(w, nil)
// }

func assets(nameDir string) http.Handler {
	fs := http.FileServer(http.Dir(nameDir))
	return http.StripPrefix("/static/", fs)
}
