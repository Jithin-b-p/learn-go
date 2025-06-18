package arrays

import "fmt"

func ArrayFunc() {

	//array declaration
	var productNames [5]string

	//declaring array of size 5 and adding only one value to it
	// var productNames = [5]string{"book"}

	// array initialization + decalaration
	prices := [5]float64{4.5, 20.5, 44.21, 34.53, 12.5}

	fmt.Println(prices)
	fmt.Println(productNames)

	//accessing data
	fmt.Println(prices[3])

	//assigning values
	productNames[2] = "pen"

	fmt.Println(productNames)

	// slicing
	// featuredPrices := prices[1:3]
	// featuredPrices := prices[:2]
	// featuredPrices := prices[2:]
	// fmt.Println(featuredPrices)

	//len and cap
	fmt.Println("len", len(prices))
	fmt.Println("cap", cap(prices))

	featuredPrices := prices[1:4]
	featuredPrices2 := featuredPrices[1:]

	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices))
	fmt.Println(cap(featuredPrices))

	fmt.Println(featuredPrices2)
	fmt.Println(len(featuredPrices2))
	fmt.Println(cap(featuredPrices2))

	fmt.Println("==============================================")

	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	s := arr[1:3]

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = s[:4]
	fmt.Println("new", s[2])

	s = append(s, 8)
	fmt.Println(s)
	fmt.Println(arr)
	fmt.Println(cap(s))
	s = append(s, 10, 11)
	fmt.Println(s)
	fmt.Println(cap(s))
	s = append(s, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	fmt.Println(s)
	fmt.Println(cap(s))

	arr1 := []int{}
	fmt.Println(cap(arr1))

	arr1 = append(arr1, 10, 20, 30)
	fmt.Println(arr1)
	fmt.Println(cap(arr1))

	fmt.Println(arr)

	fmt.Println("==============================================")

	names := []string{}

	names = append(names, "Jithin", "kiran", "kailas")

	for indx, value := range names {
		fmt.Println("Index:", indx)
		fmt.Println("Value:", value)
	}
}
