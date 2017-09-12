package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Unicorn struct {
	Name   string `bson:"name"`
	Gender string `bson:"gender"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("unicorn").C("unicorns")
	err = c.Insert(&Unicorn{"test", "this is test"})
	if err != nil {
		log.Fatal(err)
	}

	result := Unicorn{}
	err = c.Find(bson.M{"name": "Aurora"}).One(&result)
	if err != nil {
		log.Fatalf("FIND: " + err.Error())
	}

	fmt.Printf("Unicorn: %#v", result)
}
