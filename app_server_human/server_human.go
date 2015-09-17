package main

import (
	"log"
	"net/http"

	"github.com/viniciusfeitosa/rubyconf/app_server_human/handlers"
)

const (
	port string = ":8889"
)

func main() {
	http.HandleFunc("/new/human", http.HandlerFunc(handlers.PostHuman))
	http.HandleFunc("/all/humans", http.HandlerFunc(handlers.GetAllHumans))

	log.Println("Server Human start on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
