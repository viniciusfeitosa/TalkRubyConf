package main

import (
	"log"
	"net/http"

	"github.com/viniciusfeitosa/rubyconf/app_server_human/handlers"
)

func main() {

	http.HandleFunc("/new/human", http.HandlerFunc(handlers.PostHuman))
	http.HandleFunc("/all/humans", http.HandlerFunc(handlers.GetAllHumans))

	log.Println("App start on port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
