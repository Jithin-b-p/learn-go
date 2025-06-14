package main

import (
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
func main() {

	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birthdate: ")

	appUser := user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}

	// outputUserDetails(appUser)

	// passing address
	appUser.outputUserDetails()
}

// func outputUserDetails(u *user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(prompt string) (data string) {

	fmt.Print(prompt)
	fmt.Scan(&data)

	return

}
