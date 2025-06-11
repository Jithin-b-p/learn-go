package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeToFile(balance float64) {

	balanceString := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceString), 0644)

}

func readFromFile() (float64, error) {

	data, err := os.ReadFile("balance.txt")

	if err != nil {
		return 1000, errors.New("fle not found")
	}

	stringBalance := string(data)

	balance, err := strconv.ParseFloat(stringBalance, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func main() {

	accountBalance,  err := readFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------------")
	}

	fmt.Println("Welcome to the Go bank!")
	for {
		fmt.Println(`
What do you want to do?
1.Check balance
2.Deposit money
3.Withdraw money
4.Exit
	`)

		var choice int
		fmt.Print("Enter you choice: ")
		fmt.Scan(&choice)
		
		
		if choice == 1 {


			fmt.Println("Your account balance is:", accountBalance)
		
		}else if choice == 2 {
				
			var amount float64
			fmt.Print("Your amount: ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Enter a valid amount!!")
				
			}else {
				accountBalance += amount

				writeToFile(accountBalance)
				fmt.Println("Balance updated! New amount:", accountBalance)
			}
			

		}else if choice == 3 {

			var amount float64

			fmt.Print("You amount: ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Enter a valid amount!!")
				return
			}else if amount > accountBalance {
				fmt.Println("Not enough balance!!")
			}else {
				accountBalance -= amount
				writeToFile(accountBalance)
				fmt.Println("New account balance:", accountBalance)
			}
			
		}else if choice == 4 {
			fmt.Println("Thank you!! Good bye!")
			break
		}else {
			fmt.Println("Invalid Input!")
		}

	}
	
	
}