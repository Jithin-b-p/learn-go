package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Save() error {

	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	// handling error
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (note Note) Display() {
	fmt.Printf("Your note titled \"%v\" has the following content:\n\n%v\n", note.Title, note.Content)
}

// creation function
func New(title, content string) (Note, error) {

	// validation
	if title == "" || content == "" {
		return Note{}, errors.New("title and content empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
