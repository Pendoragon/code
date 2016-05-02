package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TestData struct {
	Id string `bson:"_id,omitempty"`
	Data string `bson:"data,omitempty"`
}

func main() {
	session, err := mgo.Dial("192.168.225.225")

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to mongo")
	c := session.DB("test").C("TestCollection")

	c.Upsert(bson.M{"_id": "1"}, &TestData{
		Id: "1",
		Data: "test",
	})
	defer session.Close()

	for {
		s := session.Clone()
		c := s.DB("test").C("TestCollection")

		result := &TestData{}
		err := c.Find(bson.M{"_id": "1"}).One(result)

		if err != nil {
			fmt.Printf("Got error: %+v\n", err)
		} else {
			fmt.Printf("Got result: %+v\n", *result)
		}

		time.Sleep(5 * time.Second)
		s.Close()
	}
}
