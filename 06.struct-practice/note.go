package main

import (
	"errors"
	"fmt"
)

func main() {

	title, err := getUserInput("Note title: ")
	content, err := getUserInput("Note content: ")

	if err != nil {
		panic(err)
	}

}

func getUserInput(prompt string) (string, error) {

	fmt.Println(prompt)
	var value string

	fmt.Scanln(&value)

	if value == "" {

		return "", errors.New("Value must not be empty")
	}
	return value, nil
}
