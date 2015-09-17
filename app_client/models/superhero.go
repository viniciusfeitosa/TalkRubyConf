package models

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Hero is the struct to represent a hero identity of a superhero
type Hero struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Born       string `json:"born"`
	Ocupation  string `json:"ocupation"`
	ExternalID int64  `json:"external_id"`
}

// Human is the struct to represent a human identity of a superhero
type Human struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Born       string `json:"born"`
	Ocupation  string `json:"ocupation"`
	ExternalID int64  `json:"external_id"`
}

// Superhero is the struct to represent a superhero
type Superhero struct {
	Hero Hero
	Ego  Human
}

// GetHumans is a function to get a pointer of Human struct from other service
func GetHumans() ([]Human, error) {
	var humans []Human
	resp, err := http.Get("http://localhost:8889/all/humans")
	if err != nil {
		return nil, errors.New("error to get humans json")
	}
	if err := json.NewDecoder(resp.Body).Decode(&humans); err != nil {
		return nil, errors.New("error to parse humans json")
	}
	return humans, nil
}

// GetHeros is a function to get a pointer of Hero struct from other service
func GetHeros() ([]Hero, error) {
	var heros []Hero
	resp, err := http.Get("http://localhost:8888/all/heros")
	if err != nil {
		return nil, errors.New("error to get heros json")
	}
	if err := json.NewDecoder(resp.Body).Decode(&heros); err != nil {
		return nil, errors.New("error to parse heros json")
	}
	return heros, nil
}

// GetSuperheros is a function to generate a list of superheros
func GetSuperheros() ([]*Superhero, error) {
	humans, err := GetHumans()
	if err != nil {
		return nil, errors.New("error to get humans pointer")
	}

	heros, err := GetHeros()
	if err != nil {
		return nil, errors.New("error to get heros pointer")
	}

	superheros, err := joinHumansAndHeros(humans, heros)
	if err != nil {
		return nil, errors.New("error to get superheros pointer")
	}

	return superheros, nil
}

func joinHumansAndHeros(humans []Human, heros []Hero) ([]*Superhero, error) {
	var superheros []*Superhero
	for _, human := range humans {
		for _, hero := range heros {
			if human.ExternalID == hero.ExternalID {
				superhero := &Superhero{Hero: hero, Ego: human}
				superheros = append(superheros, superhero)
			}
		}
	}

	return superheros, nil
}
