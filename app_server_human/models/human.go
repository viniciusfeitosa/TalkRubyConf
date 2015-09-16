package models

import (
	"github.com/viniciusfeitosa/rubyconf/app_server_human/configs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoSession *mgo.Session = configs.MongoSession()

// Human is the struct to represent a human identity of a superhero
type Human struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Born      string        `json:"born" bson:"born"`
	Ocupation string        `json:"ocupation" bson:"ocupation"`
}

// CreateHuman insert a new human in mongo
func CreateHuman(human Human) error {
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("RubyConf").C("human")
	if err := collection.Insert(human); err != nil {
		return err
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