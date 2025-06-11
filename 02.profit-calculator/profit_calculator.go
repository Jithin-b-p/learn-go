package main

import (
	"fmt"
	"os"
)

func saveDetails(earnings, profit, ratio float64) {

	stringData := fmt.Sprintf("Earnings: %.2f\nProfit: %.2f\nRatio: %.2f", earnings, profit, ratio)
	os.WriteFile("calculated_data.txt", []byte(stringData), 0644)

}

func main() {

	var revenue, expenses, taxRate float64 

	revenue = getUserInput("Enter the revenue: ")
	expenses = getUserInput("Enter the expenses: ")
	taxRate = getUserInput("Enter the tax-rate: ")

	earnings, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	// fmt.Printf("Earnings: %v \nProfit: %.2f \nRatio: %.2f", earnings, profit, ratio)

	saveDetails(earnings, profit, ratio)
	sEarnings := fmt.Sprintf("%.2f", earnings)
	sProfit := fmt.Sprintf("%.2f", profit)
	sRatio := fmt.Sprintf("%.2f", ratio)

	fmt.Printf("Earnings: %v\nProfit: %v\nRatio: %v\n", sEarnings, sProfit, sRatio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (earnings float64, profit float64, ratio float64) {

	earnings = revenue - expenses

	profit = earnings * (1 - taxRate / 100)

	ratio = earnings / profit

	return
}

func getUserInput(outputText string) float64 {

	var userInput float64
	fmt.Print(outputText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		panic("Invalid input can't continue..")
	}

	return userInput
}