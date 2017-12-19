package models

import (
	"github.com/go-playground/validator"
	mgo "gopkg.in/mgo.v2"
)

type Stocker struct {
	DbBase `bson:"omitempty"` // bson以降いる？
	// ID          bson.ObjectId `bson:"_id"`
	// TODO `bson:"spot_name"`のときのメンバの変数名はSpotNameではない
	Name   string  `bson:"name"`
	Amount int     `bson:"amount"`
	Price  float64 `bson:"price"`
}

func (s *Stocker) ColName() string {
	return "stocker"
}

func (s *Stocker) Insert() error {
	return s.Collection(s.ColName()).Insert(s)
}

//TODO
// func (s *Stocker) Update() error {
// 	return s.Collection(a.ColName().UpdateId(s.ID), s)
// }

func (s *Stocker) DeleteAll() (*mgo.ChangeInfo, error) {
	return s.Collection(s.ColName()).RemoveAll(nil)
}

func (s *Stocker) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
