package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jithin-b-p/learn-go/06.struct-practice/note"
	"github.com/Jithin-b-p/learn-go/06.struct-practice/todo"
)

type saver interface {
	Save() error
}

type outputter interface {
	saver
	Display()
}

func main() {

	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	text := getUserInput("Enter todo text: ")

	newNote, err := note.New(title, content)

	if err != nil {
		panic(err)
	}

	newTodo, err := todo.New(text)

	if err != nil {
		panic(err)
	}

	err = outputData(newNote)

	if err != nil {
		panic(err)
	}

	err = outputData(newTodo)

	if err != nil {
		panic(err)
	}

}

func outputData(data outputter) error {

	data.Display()
	err := saveData(data)

	return err
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving unsuccessful")
		return err
	}

	fmt.Println("saving successfully!")
	return nil
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
