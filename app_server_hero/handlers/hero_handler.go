package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/viniciusfeitosa/TalkRubyConf/app_server_hero/models"
)

func PostHero(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("It isn't a POST")
		http.Error(w, "It isn't a POST", http.StatusForbidden)
	}

	var heros []models.Hero
	if err := json.NewDecoder(r.Body).Decode(&heros); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := models.CreateHero(heros); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, hero := range heros {
		log.Printf("New hero %s created", hero.Name)
	}
	w.WriteHeader(http.StatusCreated)
}

func GetAllHeros(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("It isn't a GET")
		http.Error(w, "It isn't a GET", http.StatusForbidden)
	}

	heros, err := models.FindAllHeros()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	js, err := json.Marshal(heros)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
