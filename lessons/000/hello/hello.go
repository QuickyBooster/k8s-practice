package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// fmt.Println("Hello Hữu!!🐧")
	// fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	// message, err := greetings.Hello("Me")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
