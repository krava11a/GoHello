package main

import (
	"GoProjects/greetings"
	"fmt"
	//"rsc.io/quote"
	"log"
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"Alex","Natali","Den"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	//
	//fmt.Println("Hello and GO!!")
	//fmt.Println(quote.Go())
	//message:=greetings.Hello("Gladys")
	//fmt.Println(message)
}
