package main

import "fmt"

type test struct {
	data int
}

func main() {
	a := test{
		data: 1,
	}

	b := a
	d := &b.data
	*d = 6
	fmt.Printf("%+v\f", b)
	fmt.Printf("%+v\f", a)
}
