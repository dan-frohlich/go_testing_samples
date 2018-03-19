package main

import (
	"fmt"
	"project05/service"
)

func main() {

	f := service.NewRobotFilter()
	allowed, err := f.FilteredRobots("http://www.google.com", true)

	fmt.Printf("allowed: %+v\n", allowed)
	fmt.Printf("err: %+v\n", err)
}
