package main

import (
	"fmt"
)

// A boolean value (named after George Boole) is a special 1 bit integer type used to
// represent true and false (or on and off).
// Three logical operators are used with boolean values:
// &&	and
// ||	or
// !	not

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
