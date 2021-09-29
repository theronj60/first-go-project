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
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
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
