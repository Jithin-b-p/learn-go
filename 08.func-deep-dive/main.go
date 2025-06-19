package main

import "fmt"

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
}

//returning function (if numbers even return double else triple)
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
