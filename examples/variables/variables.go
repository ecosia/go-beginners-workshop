package main

import "fmt"

func main() {
	// First way to declare a Variable
	var myString string
	myString = "Hello"

	// Variation on first way
	var myString1 = "World"

	// Second way to declare a Variable & initiate it
	myString2 := "Hello!"

	fmt.Println(myString + " " + myString1)
	fmt.Println(myString2)

	// There are many ways to declare variables
	// and we as developers should try to keep consistency.
	// We will stick to using the first one.
	// It is also ensures the variable has it's Zero value
}
