package main

import "fmt"

func main() {
	age := 32 //regular variable

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	// adultYears := getAdultYears(agePointer)
	fmt.Println(age)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
