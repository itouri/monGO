package models

import mgo (
	"fmt"

	"gopkg.in/mgo.v2"
)

type DbBase struct {
	session  *mgo.Session
	database *mgo.Database
}

func (d *DbBase) Init() {
	d.session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(fmt.Sprintf("Initialize mongodb error:%v", err))
	}
	defer d.session.Close()
}

func (d *DbBase) Session() *mgo.Session {
	return d.session
}

func (d *DbBase) Database() *mgo.Database {
	return d.database
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