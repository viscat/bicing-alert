package app

import "gopkg.in/mgo.v2"

func GetBicingDb() mgo.Database {
	session, err := mgo.Dial("0.0.0.0") // TODO: Use environment variable instead of 0.0.0.0
	if err != nil {
		panic(err)
	}
	return *session.DB("bicing") // TODO: Use environment variable instead of bicing
}
