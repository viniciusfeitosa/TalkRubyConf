package main

import (
	"log"
	"net/http"

	"github.com/viniciusfeitosa/TalkRubyConf/app_client/handlers"
)

const (
	port string = ":8880"
)

func main() {
	superheroHandler := handlers.NewSuperheroHandler()
	http.Handle("/superheros", superheroHandler)

	log.Println("Server Superhero start on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
