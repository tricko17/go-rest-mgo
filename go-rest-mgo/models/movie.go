package models

import (
	"go-rest-mgo/db"
	"go-rest-mgo/forms"

	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:name`
	Desc   string        `json:desc`
	Rating float32       `json:rating`
}

type MovieModel struct{}

var server = "127.0.0.1"

var dbConnect = db.NewConnection(server)

func (m *MovieModel) Create(data forms.CreateMovieCommand) error {
	collection := dbConnect.Use("test-mgo", "movies")
	err := collection.Insert(bson.M{"name": data.Name, "rating": data.Rating, "desc": data.Desc})
	return err
}

func (m *MovieModel) Find() (list []Movie, err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.Find(bson.M{}).All(&list)
	return list, err
}

func (m *MovieModel) Get(id string) (movie Movie, err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MovieModel) Update(id string, data forms.UpdateMovieCommand) (err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.UpdateId(bson.ObjectIdHex(id), data)

	return err
}

func (m *MovieModel) Delete(id string) (err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.RemoveId(bson.ObjectIdHex(id))

	return err
}
