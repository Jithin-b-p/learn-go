package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// method
func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// returns the pointer so that we can mutate.
func newUser(firstName, lastName, birthDate string) (*user, error) {

	// adding validation
	if firstName == "" || lastName == "" || birthDate == "" {

		return nil, errors.New("first name, last name and birthdate are required")

	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {

	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birthdate: ")

	appUser, err := newUser(firstName, lastName, birthDate)
	// outputUserDetails(appUser)

	if err != nil {
		panic(err)
	}

	// passing address
	appUser.outputUserDetails()
}

// func outputUserDetails(u *user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(prompt string) (data string) {

	fmt.Print(prompt)
	fmt.Scanln(&data)

	return

}
