package models

import (
	"github.com/viniciusfeitosa/TalkRubyConf/app_server_hero/configs"
	"gopkg.in/mgo.v2/bson"
)

var mongoSession = configs.MongoSession()

// Hero is the struct to represent a hero identity of a superhero
type Hero struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Born       string        `json:"born" bson:"born"`
	Ocupation  string        `json:"ocupation" bson:"ocupation"`
	ExternalID int64         `json:"external_id" bson:"external_id"`
}

// CreateHero insert a new hero in mongo
func CreateHero(heros []Hero) error {
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("RubyConf").C("hero")
	for _, hero := range heros {
		if err := collection.Insert(hero); err != nil {
			return err
		}
	}
	return nil
}

func FindAllHeros() ([]Hero, error) {
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var heros []Hero
	collection := sessionCopy.DB("RubyConf").C("hero")
	if err := collection.Find(nil).All(&heros); err != nil {
		return heros, err
	}

	return heros, nil
}
