package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	//If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//Hellos returns a map that associates each of the named people
//with a gretting messages.

func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages.
	messages := make(map[string]string)
	// loop through the received slice of names, calling
	// the hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Will met!",
	}
	return formats[rand.Intn(len(formats))]
}
