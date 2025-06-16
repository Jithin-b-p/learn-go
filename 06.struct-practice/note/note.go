package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v", note.title, note.content)
}

// creation function
func New(title, content string) (Note, error) {

	// validation
	if title == "" || content == "" {
		return Note{}, errors.New("title and content empty")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
