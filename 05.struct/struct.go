package main

import (
	"fmt"
	"github.com/Jithin-b-p/learn-go/05.struct/user"
)

func main() {

	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birthdate: ")

	appUser, err := user.New(firstName, lastName, birthDate)
	// outputUserDetails(appUser)

	if err != nil {
		panic(err)
	}

	// passing address
	appUser.OutputUserDetails()
}

// func outputUserDetails(u *user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(prompt string) (data string) {

	fmt.Print(prompt)
	fmt.Scanln(&data)

	return

}
