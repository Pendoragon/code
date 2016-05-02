package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	for {
		s := session.Clone()
		err := s.Ping()

		if err != nil {
			fmt.Printf("Got error: %+v", err)
		}

		time.Sleep(5 * time.Second)
		s.Close()
	}
}
