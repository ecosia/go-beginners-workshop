package main

import (
	"fmt"
)

// Numbers are vital for programming and there are a few different types you can use, below we cover the main
// types of numbers and some simple maths.

var (
	i int     // Integers – like their mathematical counterpart – are numbers without a decimal component.
	f float64 // Floating point numbers are numbers that contain a decimal component (real numbers).
)

func main() {
	i := 42
	fmt.Printf("i is of type %T\n", i)
	f := float64(i) // This is called type conversion where we tale a var with one type and convert it to another.
	fmt.Printf("f is of type %T\n", f)

	var zeroValue int // when we declare a variable but dont give it a value it takes the zero value for a int this is 0
	fmt.Printf("zeroValue has a value of %d\n", zeroValue)

	// There are different sizes of int
	// int  int8  int16  int32  int64
	// The number after the word int describes the size of the integer. For the most part you can just use int.

	// Run the code below to see how these math statements work, then try some of your own! :)
	var (
		a int
		b int
		c float64
		d float64
	)
	a = 5
	b = 8
	c = 2.00
	d = 5.50

	fmt.Println(a + b)
	fmt.Println(d - c)
	fmt.Println(c * d)
	fmt.Println(b / 2)
}
