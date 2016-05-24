package main

import "fmt"

func main() {
	test()
	select{}

}

func test() {
	go func() {
		for {
			fmt.Println("echo")
		}
	}()
	fmt.Println("out")
}
