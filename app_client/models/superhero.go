package models

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
	HeroName       string `json:"hero_name"`
	HumanName      string `json:"human_name"`
	HeroBorn       string `json:"hero_born"`
	HumanBorn      string `json:"human_born"`
	HeroOcupation  string `json:"hero_ocupation"`
	HumanOcupation string `json:"human_ocupation"`
}
