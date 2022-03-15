package api

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

type Value struct {
	Id      string `json:"id" bson:"_id,omitempty"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func init() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("Failure connect to %v", err)
	}
	db = session.DB("Values")
}

func collection() *mgo.Collection {
	return db.C("values")
}

func GetAll() ([]Value, error) {
	res := []Value{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func Get(id string) (*Value, error) {
	res := Value{}

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("invalid id")
	}

	if err := collection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func Create(value Value) error {
	return collection().Insert(value)
}

func Delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid id")
	}
	return collection().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
