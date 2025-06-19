package variadicfn

import "fmt"

func Sum(startingValue int, nums ...int) {

	sum := 0
	for _, value := range nums {
		sum += value
	}
	fmt.Println("Starting value:", startingValue)
	fmt.Println("sum is:", sum)
}
