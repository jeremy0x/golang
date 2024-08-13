package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
    /* 
			set properties of the predefined Logger, including
		  the log entry prefix and a flag to disable printing
		  the time, source file, and line number.
		*/
		log.SetPrefix("greetings: ")
		log.SetFlags(0)

		// a slice of names
		names := []string{"Jeremy", "Jenny", "Jasmine"}

		// request greeting messages for the names
		messages, error := greetings.Hellos(names)
		if error != nil {
			log.Fatal(error)
		}

		// if no error was returned, print the returned map of messages
		fmt.Println(messages)
}
