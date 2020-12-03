package main

import (
	"fmt"
)

func main() {
	fmt.Println("ArraySlicesMaps!")
	// declaring an array of int with size 5 - valid usable indices are obviously from 0 to 4
	var ara [5]int
	// assigning values in some indices manually, others are initialized as zero by default
	ara[0] = 1
	ara[1] = 3
	ara[4] = 5
	// print the whole array like a list of int
	fmt.Println(ara)
	// printing value at a particular position, 0 in this case
	fmt.Println(ara[0])

	// looping through the array using it's size of lenght using the len() func.
	for i := 0; i < len(ara); i++ {
		fmt.Print(ara[i], " ")
	}
	fmt.Println()

	// iterating over the array, using the range() of index, value pair, remeber both index, value must be used as they are declared
	for index, value := range ara {
		fmt.Println(index, value)
	}
	// if we don't want to use the index, we can simply replace it with an underscore, remember underscore(_) can't hold a value
	for _, value := range ara {
		fmt.Print(value, " ")
	}
	fmt.Println()
}
