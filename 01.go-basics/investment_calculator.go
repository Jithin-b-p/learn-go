package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5

	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Enter the investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	realFutureValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(realFutureValue)

}