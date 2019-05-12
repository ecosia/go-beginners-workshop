package main

import (
	"fmt"
)

// Loops are an integral part of programming!
// Go has only one looping construct, the for loop.
// Read the code below and see if you understand what is going on.
// When you run the code, it will fail, are you able to make it run? :)

func main() {
	sum := 0

	// The for Loop can take a start point, end point (condition) and interval.
	// When the finish condition is met we exit out of the loop
	// and continue to the next part of the code.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// We can also write it with only the condition
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	// In other languages there is also a while Loop. In GO we use
	// a for loop to implement this logic (Note we can also drop the semi-colons)
	for sum < 1000 {
		sum += sum
	}
	// Note that the printed SUM is above 100 that is because when we
	// check the condition SUM is only 90 and therefore less than 100
	// therefore our loop code runs, only on the next check is the condition
	// false and we exit the loop!
	fmt.Println(sum)

	// BEWARE of infinite loops! if your condition never stops being TRUE the loop
	// will run forever and break your code. Alter this loop so it does not exit
	// and prints the rest of our code :)
	for {
		sum += sum
	}

	fmt.Println("Congratulations you made the code exit the loop and print this line")

	// As a added bonus here is a Loop which goes through a range of items (in this case strings)
	// within a array ( we will cover arrays later so do not worry if you are not familiar with them!
	// This is a little more complex because it has an index as well as a value, don't worry if you don't
	// fully understand this yet.
	strings := []string{"hello", "world"}
	for index, string := range strings {
		fmt.Println(index, string)
	}
}
