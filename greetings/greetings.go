package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("no name provided")
	}

	// create a message using a random format
	message := fmt.Sprintf(randomFormat(), name)

	// use the below line to test the error handling
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

// hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)

	// loop through the names and generate a message for each
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}

		// associate the message with the name
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random.
// note: the name of the function starts with a lowercase letter, making it private to the package
func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Yo, %v! What's up?",
		"Hey, %v! How's it going?",
		"Hello, %v! How are you?",
		"Hiya, %v! What's new?",
		"Hey there, %v! How's life?",
	}

	// return a randomly selected message format by specifying a random index
	return formats[rand.Intn(len(formats))]
}
