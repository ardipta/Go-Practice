package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// this function return a string type value and
// the parameter is also string type
func Hello(name string) (string, error) {
	// if no name, then return an empty message
	if name == "" {
		return "", errors.New("Empty Name")
	}
	//if there is name then it return the name
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well Met",
	}

	return formats[rand.Intn(len(formats))]
}
