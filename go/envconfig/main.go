package main

import (
    "fmt"
    "log"
    "time"

    "github.com/kelseyhightower/envconfig"
)

type Specification struct {
	Debug   bool `default:"true"`
	Port    int `default:"80"`
	User    string `default:"pen//.."`
	Users   []string `default:"a,b,c"`
	Rate    float32 `default:"0.9"`
	Timeout time.Duration `default:"2m" envconfig:"time_out"`
}

func main() {
	var s Specification
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Debug: %v\nPort: %d\nUser: %s\nRate: %f\nTimeout: %s\n"
	_, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Rate, s.Timeout)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf("  %s\n", u)
	}
}
