package configs

import (
	"log"

	"gopkg.in/mgo.v2"
)

type MongoDB interface {
	MongoSession() *mgo.Session
}

// MongoSession return a session with mongo
func MongoSession() *mgo.Session {
	mongoSession, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal("Create session: %s\n", err.Error())
	}

	return mongoSession
}
