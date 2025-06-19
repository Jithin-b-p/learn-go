package main

import (
	"fmt"

	"github.com/Jithin-b-p/learn-go/08.func-deep-dive/variadicfn"
)

type transformFn func(int) int

func main() {

	nums := []int{1, 2, 4, 6}

	dNums := transformNumber(&nums, doubleValue)
	tNums := transformNumber(&nums, tripleValue)

	fmt.Println(dNums)
	fmt.Println(tNums)

	evenNums := []int{2, 4, 6, 8}
	oddNums := []int{3, 5, 7, 9}

	transformedNums := transformNumber(&evenNums, getTransformFunc(&evenNums))
	fmt.Println(transformedNums)

	transformedNums = transformNumber(&oddNums, getTransformFunc(&oddNums))
	fmt.Println(transformedNums)

	//anonymous functions
	sNums := transformNumber(&nums, func(num int) int {
		return num * num
	})

	fmt.Println("square of nums", sNums)

	doubled := transformNumber(&nums, createTransformer(2))
	fmt.Println("double created transformer:", doubled)

	//find factorial

	fmt.Println("factorial of 5 is", factorial(5))

	fmt.Println("=========================================")

	variadicfn.Sum(20, 4, 8, 16, 32, 64)

	//unpacking slice
	variadicfn.Sum(10, nums...)

}

// recursion
func factorial(n int) int {

	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

// returning function (if numbers even return double else triple)
func getTransformFunc(nums *[]int) transformFn {

	if (*nums)[0]%2 == 0 {
		return doubleValue
	}

	return tripleValue
}

func transformNumber(nums *[]int, transform transformFn) []int {

	tNumbers := []int{}
	for _, value := range *nums {
		tNumbers = append(tNumbers, transform(value))
	}

	return tNumbers
}
func doubleValue(value int) int {
	return 2 * value
}
func tripleValue(value int) int {
	return 3 * value
}

// we are creating a factory function
func createTransformer(factor int) func(num int) int {

	return func(num int) int {
		return num * factor
	}
}
