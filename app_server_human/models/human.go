package models

import (
	"github.com/viniciusfeitosa/TalkRubyConf/app_server_human/configs"
	"gopkg.in/mgo.v2/bson"
)

var mongoSession = configs.MongoSession()

// Human is the struct to represent a human identity of a superhero
type Human struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Born       string        `json:"born" bson:"born"`
	Ocupation  string        `json:"ocupation" bson:"ocupation"`
	ExternalID int64         `json:"external_id" bson:"external_id"`
}

// CreateHuman insert a new human in mongo
func CreateHuman(humans []Human) error {
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("RubyConf").C("human")
	for _, human := range humans {
		if err := collection.Insert(human); err != nil {
			return err
		}
	}
	return nil
}

func FindAllHumans() ([]Human, error) {
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var humans []Human
	collection := sessionCopy.DB("RubyConf").C("human")
	if err := collection.Find(nil).All(&humans); err != nil {
		return humans, err
	}

	return humans, nil
}
