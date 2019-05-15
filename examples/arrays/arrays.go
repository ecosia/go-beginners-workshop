package main

import (
	"fmt"
)

func main() {
	// Declare an array of 5 strings
	var movies [2]string
	fmt.Println(movies)

	// Initialize the array
	movies[0] = "Interstellar"
	movies[1] = "The eternal sunshine of the spotless mind"
	fmt.Println(movies[0])
	fmt.Println(movies[1])

	// Now you can replace one of the movies with your favorite movie
	// and print the result
	// favoriteMovie := ""
	// movies....
	// fmt.Println(..)

	// Declare and initialize an array of 3 strings
	colors := [3]string{"black", "blue", "purple"}
	fmt.Println("Colors: ", colors)

	// Declare and initializes an array of 2 strings
	newColors := [2]string{"pink", "green"}
	fmt.Println("newColors: ", newColors)

	// Copy values from newColors array to colors array.
	// What happens when uncommenting the next line?

	// colors = newColors

	// Declare and initialize an integer array. The size of the array will be
	// determined based on the number of initialized values
	numbers := [...]int{10, 20, 30}
	fmt.Println(numbers)

	// Initialize result
	result := 0

	// Loop through numbers array and sum the values
	for i, number := range numbers {
		// Show number value and its address
		fmt.Println(number, &numbers[i])
		// Adds each value to the result variable
		result = result + number
	}
	fmt.Println("Result: ", result)
}
