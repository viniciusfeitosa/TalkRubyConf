package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/viniciusfeitosa/TalkRubyConf/app_client/models"
)

// SuperheroHandler is the handler to get the superhero list
type SuperheroHandler struct {
}

// NewSuperheroHandler generate a new instance of SuperheroHandler
func NewSuperheroHandler() *SuperheroHandler {
	return &SuperheroHandler{}
}

func (s *SuperheroHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	superheros, err := models.GetSuperheros()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error when got superheros", http.StatusInternalServerError)
	}

	js, err := json.Marshal(superheros)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
