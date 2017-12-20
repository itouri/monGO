package models

import (
	"github.com/go-playground/validator"
	mgo "gopkg.in/mgo.v2"
)

type Sell struct {
	DbBase `bson:"omitempty"` // bson以降いる？
	// ID          bson.ObjectId `bson:"_id"`
	// TODO `bson:"spot_name"`のときのメンバの変数名はSpotNameではない
	Sell float64 `bson:"sell"`
}

func (s *Sell) ColName() string {
	return "sell"
}

func (s *Sell) Insert() error {
	return s.Collection(s.ColName()).Insert(s)
}

func (s *Sell) Upsert(selector, upsert interface{}) (*mgo.ChangeInfo, error) {
	return s.Collection(s.ColName()).Upsert(selector, upsert)
}

func (s *Sell) DeleteAll() (*mgo.ChangeInfo, error) {
	return s.Collection(s.ColName()).RemoveAll(nil)
}

func (s *Sell) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
