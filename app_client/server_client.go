package main

import (
	"log"
	"net/http"

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
	mux.Get("/static/", assets("static"))

	log.Println("Server Superhero start on port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func assets(nameDir string) http.Handler {
	fs := http.FileServer(http.Dir(nameDir))
	return http.StripPrefix("/static/", fs)
}
