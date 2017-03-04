package main

import "gopkg.in/mgo.v2"

var db *mgo.Database

func init() {
	initDb()

}

func initDb() {
	session, err := mgo.Dial("0.0.0.0")
	if err != nil {
		panic(err)
	}
	db = session.DB("bicing")
}
