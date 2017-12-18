package models

import (
	"fmt"
	"log"

	"../config"

	"gopkg.in/mgo.v2"
)

var (
	_Session  *mgo.Session
	_Database *mgo.Database
)

func init() {
	_Session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", confing.DBCONFIG.Host, confing.DBCONFIG.Port))
	if err != nil {
		panic(fmt.Sprintf("Initialize mongodb error:%v", err))
	}
	_Database = _Session.DB(confing.DBCONFIG.Database)
	if err = _Session.Ping(); err != nil {
		panic(fmt.Sprintf("MongoDB execute ping error:%v", err))
	}
	log.Println("MongoDB initialize success.")
	mgo.SetDebug(confing.APPCONFIG.Debug)
}
