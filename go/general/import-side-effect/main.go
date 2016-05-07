package main

import (
	"time"
	provider "github.com/Pendoragon/code/go/general/import-side-effect/provider"
	_ "github.com/Pendoragon/code/go/general/import-side-effect/provider/providers/aws"
)

func main() {
	provider.Print()
	time.Sleep(3 * time.Second)
}
