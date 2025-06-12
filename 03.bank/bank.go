package main

import (
	"fmt"

	"github.com/Jithin-b-p/learn-go/03.bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const fileName = "balance.txt"

func main() {

	accountBalance, err := fileops.ReadFloatFromFile(fileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------------")
	}

	fmt.Println("Welcome to the Go bank!")
	fmt.Println(`avaliable 24/7`, "contact:", randomdata.PhoneNumber())
	for {

		presentOptions()

		var choice int
		fmt.Print("Enter you choice: ")
		fmt.Scan(&choice)

		if choice == 1 {

			fmt.Println("Your account balance is:", accountBalance)

		} else if choice == 2 {

			var amount float64
			fmt.Print("Your amount: ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Enter a valid amount!!")

			} else {
				accountBalance += amount

				fileops.WriteFloatToFile(accountBalance, fileName)
				fmt.Println("Balance updated! New amount:", accountBalance)
			}

		} else if choice == 3 {

			var amount float64

			fmt.Print("You amount: ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Enter a valid amount!!")
				return
			} else if amount > accountBalance {
				fmt.Println("Not enough balance!!")
			} else {
				accountBalance -= amount
				fileops.WriteFloatToFile(accountBalance, fileName)
				fmt.Println("New account balance:", accountBalance)
			}

		} else if choice == 4 {
			fmt.Println("Thank you!! Good bye!")
			break
		} else {
			fmt.Println("Invalid Input!")
		}

	}

}
