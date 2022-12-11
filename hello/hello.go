package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// get the greeting message and print it
	message, err := greetings.Hello("Ashiq")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
