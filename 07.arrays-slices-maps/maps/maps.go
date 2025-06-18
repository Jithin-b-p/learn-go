package maps

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

func MapFunc() {

	//creating an empty map
	// websites := map[string]string{}

	websites := map[string]string{

		"google": "www.google.com",
		"aws":    "www.aws.com",
	}

	fmt.Println(websites)

	fmt.Println(websites["google"])

	websites["linkedin"] = "www.linkedin.com"

	fmt.Println(websites)

	delete(websites, "linkedin")
	fmt.Println(websites)

	fmt.Println("========================================")

	numbers := make([]int, 0, 5)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers, 1, 2, 4, 31, 12, 11)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	fmt.Println("========================================")

	// itemPrices := make(map[string]float64, 5)

	// fmt.Println(itemPrices)
	// fmt.Println(len(itemPrices))

	fmt.Println("========================================")
	itemPrices := make(floatMap, 5)

	itemPrices["pen"] = 10.4
	itemPrices["pencil"] = 5.3

	itemPrices.output()

	fmt.Println("========================================")

	for key, value := range itemPrices {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
