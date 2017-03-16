package app

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Email string
	Alerts []Alert
}


type UserRepository struct {
	Db mgo.Database
}

func (u UserRepository) GetUser(email string) (User, error) {
	user := User{}
	c := u.Db.C("user")
	err := c.Find(bson.M{}).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u UserRepository) New(email string) (User, error) {
	user := User{Email: email}
	c := u.Db.C("user")
	err := c.Insert(user)
	if err != nil {
		return user, err
	}
	return user, nil
}