package greetings

import (
	"errors"
	"fmt"
)

// this function return a string type value and
// the parameter is also string type
func Hello(name string) (string, error) {
	// if no name, then return an empty message
	if name == "" {
		return "", errors.New("Empty Name")
	}
	//if there is name then it return the name
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
