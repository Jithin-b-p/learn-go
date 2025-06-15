package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// method
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// returns the pointer so that we can mutate.
func New(firstName, lastName, birthDate string) (*User, error) {

	// adding validation
	if firstName == "" || lastName == "" || birthDate == "" {

		return nil, errors.New("first name, last name and birthdate are required")

	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
