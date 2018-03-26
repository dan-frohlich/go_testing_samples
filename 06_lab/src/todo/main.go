package main

import (
	"fmt"
	"log"
)

func main() {
	// Open the data.db file. It will be created if it doesn't exist.
	err := initializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer closeDB()

	addToDo("write a test")
	addToDo("run a test")
	addToDo("write code to make it pass")
	addToDo("refactor code")

	for _, item := range todo() {
		fmt.Println(item)
	}

	markDone("write a test")

	fmt.Println("-----------------")
	for _, item := range todo() {
		fmt.Println(item)
	}

}
