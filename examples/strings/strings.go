package main

import (
	"fmt"
	"strconv"
)

var (
	firstName   string
	secondName  string
	aVerb       string // ie jump/run/sleep
	luckyNumber int
)

func main() {
	// Below initialise the variables above with values
	// An example would be firstName = Jessica

	// Printing Strings - click run to print the strings
	fmt.Println("Hi there", firstName, secondName, "!")

	// Multiline String you will see this also printed when you click run!
	var multiLineString = `We are super happy
to hear you want to learn GO!`

	fmt.Println(multiLineString)

	// A different way to print strings (uncomment the lines below)
	// fmt.Printf("How many Developers does it take to make a Gopher %v?\n", aVerb)

	// convert another data type to string (uncomment the lines below)
	// var number string
	// number = strconv.Itoa(luckyNumber) // this converts an Int to a string

	// Joining strings and not printing
	// var sentence string
	// sentence = fmt.Sprintf("It takes %s Developers to make a Gopher %s ^^ just kidding, you can do it!", number, aVerb)

	// Now put the code to print the sentence we created above and click run :)
}
