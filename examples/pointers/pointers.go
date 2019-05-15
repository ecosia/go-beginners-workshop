package main

import (
	"fmt"
)

var (
	number int
)

func multiply(i int) int {
	i *= 10
	fmt.Println("number value inside multiply func: ", i)
	return i
}

func reset(i *int) {
	*i = 0
}

func main() {
	// When we work with values there are two ways we can pass them, say to a function.
	// The first is what we call "pass by value"
	number = 5
	multiply(number)
	// When we pass a value in this way the function receives it's own copy of number
	// This means if the multiply function alters the number passed to it, it only affects
	// its copy and not the original variable.
	fmt.Println("number value outside multiply func: ", number)

	// When we pass by value we are creating copies each time of our value and this can
	// be a problem for memory, particular if we have low resources (such as on embedded systems for IoT)
	// The other issue is we may want to alter a variable and be able to access it with the changes later.

	// The second way we can pass a value is by passing its memory address - we call this a pointer.
	// It uses the & to indicate it is a pointer. However you will see in the function signature we use a
	// *
	reset(&number)
	fmt.Println("number value after reset func: ", number)

	// Pointers can be printed too.
	fmt.Println("pointer:", &number)

}
