package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/unrolled/render"
	"github.com/viniciusfeitosa/TalkRubyConf/app_client/models"
)

// SuperheroHandler is the handler to get the superhero list
type SuperheroHandler struct {
	render *render.Render
	Lista  []*models.Superhero
}

// NewSuperheroHandler generate a new instance of SuperheroHandler
func NewSuperheroHandler(render *render.Render) *SuperheroHandler {
	return &SuperheroHandler{render: render}
}

func (s *SuperheroHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	superheros, err := models.GetSuperheros()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error when got superheros", http.StatusInternalServerError)
	}

	s.Lista = superheros

	s.render.HTML(w, http.StatusOK, "index", s)
}

// JSON is as API with de content in json
func (s *SuperheroHandler) JSON(w http.ResponseWriter, r *http.Request) {
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
