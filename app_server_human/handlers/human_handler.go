package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/viniciusfeitosa/rubyconf/app_server_human/models"
)

func PostHuman(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("It isn't a POST")
		http.Error(w, "It isn't a POST", http.StatusForbidden)
	}

	var humans []models.Human
	if err := json.NewDecoder(r.Body).Decode(&humans); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := models.CreateHuman(humans); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, human := range humans {
		log.Printf("New human %s created", human.Name)
	}
	w.WriteHeader(http.StatusCreated)
}

func GetAllHumans(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("It isn't a GET")
		http.Error(w, "It isn't a GET", http.StatusForbidden)
	}

	humans, err := models.FindAllHumans()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	js, err := json.Marshal(humans)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
