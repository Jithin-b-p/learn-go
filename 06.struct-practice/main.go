package main

import (
	"fmt"

	"github.com/Jithin-b-p/learn-go/06.struct-practice/note"
)

func main() {

	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	newNote, err := note.New(title, content)

	if err != nil {
		panic(err)
	}

}

func getUserInput(prompt string) string {

	fmt.Println(prompt)
	var value string

	fmt.Scanln(&value)
	return value
}
