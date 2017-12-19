package models

import (
	"github.com/go-playground/validator"
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

//TODO
// func (s *Stocker) Delete() error {
// 	return s.Collection(s.ColName().RemoveId(s.ID))
// }

func (s *Stocker) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
