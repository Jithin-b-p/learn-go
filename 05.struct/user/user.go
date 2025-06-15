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

type Admin struct {
	email    string
	password string
	User
}

// method
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func NewAdmin(firstName, lastName, birthDate, email, password string) *Admin {

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: firstName,
			lastName:  lastName,
			birthDate: birthDate,
		},
	}
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
