package models

import (
	"time"

	"github.com/go-playground/validator"
)

type Spot struct {
	DbBase `bson:"omitempty"` // bson以降いる？
	// ID          bson.ObjectId `bson:"_id"`
	// TODO `bson:"spot_name"`のときのメンバの変数名はSpotNameではない
	Spot_name   string    `bson:"spot_name"`
	User_id     int       `bson:"user_id"`
	Image_ids   []int     `bson:"image_ids"`
	Latitude    float64   `bson:"latitude"`
	Longitude   float64   `bson:"longitude"`
	Prefecture  string    `bson:"prefecture"`
	City        string    `bson:"city"`
	Description string    `bson:"description"`
	Hint        string    `bson:"hint"`
	Favorites   int       `bson:"favorites"`
	Created     time.Time `bson:"created"`
	Modified    time.Time `bson:"modified"`
}

func (s *Spot) ColName() string {
	return "spot"
}

func (s *Spot) Insert() error {
	return s.Collection(s.ColName()).Insert(s)
}

//TODO
// func (s *Spot) Update() error {
// 	return s.Collection(a.ColName().UpdateId(s.ID), s)
// }

//TODO
// func (s *Spot) Delete() error {
// 	return s.Collection(s.ColName().RemoveId(s.ID))
// }

func (s *Spot) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
