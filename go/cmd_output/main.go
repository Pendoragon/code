package main

import (
	// "encoding/json"
	"fmt"
	// "io"
	"log"
	"os/exec"
	// "bufio"
	// "bytes"
)

func main() {
	cmd := exec.Command("bash", "test.sh")
	stdout, err := cmd.StdoutPipe()
	// stderr, _ := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	b := make([]byte, 100)
	_, err = stdout.Read(b)
	for err == nil {
		fmt.Printf("++++++++++" + string(b) + "\n")
		_, err = stdout.Read(b)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
