package main

import (
	"fmt"
)

type Status string

const (
	Preparing     Status = "preparing"
)

func main(){
	str := "preparing"
	status := Status(str)

	fmt.Println(status)
}
