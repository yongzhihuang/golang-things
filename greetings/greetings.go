package greetings

import (
	"errors"
	"fmt"
)

// Hello is capital, this means it's exported
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
