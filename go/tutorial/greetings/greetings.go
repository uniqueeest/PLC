package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)

    // Loop through the received slice of names, calling the Hello function
    // for each name to get a message.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }
    return messages, nil
}

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format.
    return formats[randomGenerator.Intn(len(formats))]
}