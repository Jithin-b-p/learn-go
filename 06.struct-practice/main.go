package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jithin-b-p/learn-go/06.struct-practice/note"
)

func main() {

	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	newNote, err := note.New(title, content)

	if err != nil {
		panic(err)
	}
	newNote.Display()

	err = newNote.Save()

	if err != nil {
		panic(err)
	}
}

func getUserInput(prompt string) string {

	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	data, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	data = strings.TrimSuffix(data, "\n")
	data = strings.TrimSuffix(data, "\r")

	return data
}
