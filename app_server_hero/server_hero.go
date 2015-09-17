package main

import (
	"log"
	"net/http"

	"github.com/viniciusfeitosa/rubyconf/app_server_hero/handlers"
)

const (
	port string = ":8888"
)

func main() {
	http.HandleFunc("/new/hero", http.HandlerFunc(handlers.PostHero))
	http.HandleFunc("/all/heros", http.HandlerFunc(handlers.GetAllHeros))

	log.Println("Server Hero start on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
