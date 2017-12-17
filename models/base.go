package models

import mgo (
	"fmt"

	"gopkg.in/mgo.v2"
)

type DbBase struct {}

func (d *DbBase) Session() *mgo.Session {
	return _Session
}

func (d *DbBase) Database() *mgo.Database {
	return _Database
}

func (d *DbBase) Collection(collectionName string) *mgo.Collection {
	return d.Database().C(collectionName)
}

func (d *DbBase) Find(collectionName string, query, selector interface{}) *mgo.Query {
	return d.Database().C(collectionName).Find(query).Select(selector)
}

func (d *DbBase) FindId(collectionName string, id int) *mgo.Query {
	return d.Database().C(collectionName).FindId(id)
}