package main

import "fmt"

func main() {

	age := 32 //regular variable

	// var agePointer *int = &age
	agePointer := &age //shortcut pointer variable.

	fmt.Println(agePointer)  // This will gives us the address it holds
	fmt.Println(*agePointer) // dereferencing it gives the value that the address holds.

	adultYears := getAdultYears(&age)

	fmt.Println(adultYears)

	fmt.Println("===============================")

	fmt.Println("Age before mutation:", age)
	mutateAge(&age)
	fmt.Println("Age after mutation:", age)
}

func mutateAge(age *int) {

	*age = *age - 5
}

func getAdultYears(age *int) int {

	fmt.Println(age) //same address as that of age in main function
	return *age - 18
}
