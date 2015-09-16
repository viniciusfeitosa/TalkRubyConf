package models

import "gopkg.in/mgo.v2/bson"

// Hero	is the struct to represent a human identity of a superhero
type Hero struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	Born      string        `bson:"born"`
	Ocupation string        `bson:"ocupation"`
}
