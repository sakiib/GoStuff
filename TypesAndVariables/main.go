package main

import (
	"fmt"
)

func main() {
	fmt.Println("Types & Variables")

	// this is how to add a single line comment
	// if anything is declared, then it must be used

	// declaring variables in different ways, declarring first, then assigning value, default is always zero
	var i int
	i = 5

	// how to read: we're declaring a variable j which is an int type variable
	// we're declaring a variable & assigning the value at the same time
	var j int = 10

	// if we don't mention the type, then it'll see the value & assign the type, string for this instance
	var name = "sakib"
	fmt.Println(i, j, name)

	// shorthand for declarign & assigning values to a new variable
	k := 15
	// we can't re-declare it like this: k := 10, what we can do is assign a new value to it like this: k = 10

	// Printf - value type newline, '\n' is for the newline as usual, %v = value, %T = type, should print: 5, int
	fmt.Printf("%v %T\n", i, i)

	// Print without & with a new line, takes one or more comma separated arguments
	fmt.Print(i, j, k)
	fmt.Println(i, j, k)
}
