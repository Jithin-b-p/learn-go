package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
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
