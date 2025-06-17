package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Save() error {

	fileName := "todo.json"

	json, err := json.Marshal(todo)

	// handling error
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

// creation function
func New(content string) (Todo, error) {

	// validation
	if content == "" {
		return Todo{}, errors.New("title and content empty")
	}
	return Todo{
		Text: content,
	}, nil
}
