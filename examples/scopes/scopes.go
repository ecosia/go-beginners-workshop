package main

import (
	"fmt"
)

// Scope is an important concept to understand for any programming language
// it refers to where a variable is available / accessible.
// If an variable is not accesible in a location (within a scope) undefined will be returned
// when we try access it.
// Read the code below and see if you understand this concept, when you run it you will recieve an
// error saying innerLoopCat is undefined. Try alter the code
// so it will run, there are a few ways to do this so it's up to you to
// see what works :)

var (
	globalCat = "Henry"
	// A global variable such as this, is declared at the highest level and accessible
	// anywhere in the package.
	ExportableGlobalCat = "Vanessa"
	// If the variable has a capital letter at the start this means it can be exported and
	// if we import the package else where we can acess it from other packages.
	// For this workshop we will not cover this but you can read more about it here:
	//
)

func main() {
	var localCat string
	localCat = "James"
	// This variable is only accessible within the main function but can be used anywhere within it.

	// Do not worry if you don't understand the loop below, we will cover loops in a separate section.
	for i := 1; i <= 2; i++ {
		var innerLoopCat string
		innerLoopCat = "Mellisa"
		fmt.Println(localCat)
		fmt.Println(innerLoopCat)
		fmt.Println(globalCat)
	}
	fmt.Println(innerLoopCat)
	fmt.Println(globalCat)
}
