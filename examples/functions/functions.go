package main

import (
	"fmt"
)

// In programming functions are fundamental!

// we declare a function with the keyword func
// as well as the function's name (in this example we created 2 functions: myFunction and printString)
// we also need to declare what arguments the function takes. In this case, a variable 's'.
// As Go is a statically typed language we also need to include what type 's' is.
// Lastly we add to the function signature (what the function returns) - in this case a string!

func myFunction(s string) string {
	return fmt.Sprintf("%s are magical", s)
}

func printString(s string) {
	fmt.Println(s)
}

var (
	yourString string
)

// Every Go program has at least one function, which is main().
func main() {
	// we call a function by writing it's name with round brackets after it.
	printString("Hello Gopher!")
	// give the variable yourString a value below, then press run

	// If the function takes arguments we add them inside the brackets.
	// If the function returns something we need to make sure we declare a variable and assign it.
	v := myFunction(yourString)
	printString(v)
}
