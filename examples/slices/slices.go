package main

import (
	"fmt"
)

func main() {
	// Create a slice of strings
	days := []string{"pondělí", "úterý", "středa", "čtvrtek", "pátek"}
	fmt.Println("Days: ", days)

	// Append strings to the end of the slice
	days = append(days, "sobota", "neděle")
	fmt.Println("Days: ", days)

	// Take the last appended days, that are the ones on index 5 and 6
	weekend := days[5:7]
	fmt.Println("Weekend: ", weekend)

	// Declare an integer slice
	var numbers []int
	fmt.Println("Numbers: ", numbers)

	// Populate the slice
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i*i)
	}
	fmt.Println("Numbers: ", numbers)

}
