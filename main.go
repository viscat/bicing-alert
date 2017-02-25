package main

import (
	"fmt"
	"bicing-alert/bicing"
	"gopkg.in/mgo.v2"
	"log"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	status, err:= bicing.GetStationsStatus()
	if err != nil {
		panic(err)
	}

	fmt.Printf("update time: %v", status)

	session, err := mgo.Dial("0.0.0.0")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("bicing").C("status")
	err = c.Insert(status)
	if err != nil {
		log.Fatal(err)
	}

}


