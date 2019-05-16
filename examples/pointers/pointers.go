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
	// When we work with parameter values there are two ways we can pass them to a function.
	// The first is what we call "pass by value"
	number = 5
	multiply(number)
	// When we pass a value in this way the function receives its own *copy* of the number variable
	// This means if the multiply function alters the number passed to it, it only affects
	// its copy and not the original variable.
	fmt.Println("number value outside multiply func: ", number)

	// When we pass a parameter "by value" we are creating, each time, a copy of our original variable. This can
	// be a problem for the memory, in particular if we have low resources (such as on embedded systems for IoT).
	// The other issue is that we may want to alter a variable and be able to access it with the changes later.

	// The second way we can pass a parameter value is by passing its memory address - we call this a reference.
	// We use the & operator to create a "pointer" to our number variable. A pointer is a special type that stores a memory address
	// for a value of some specific type. The pointer declaration is marked with a * followed by the pointed type
	// (as seen on the reset method parameter definition)
	reset(&number)
	fmt.Println("number value after reset func: ", number)

	// Pointers can be printed too.
	fmt.Println("pointer:", &number)

}
