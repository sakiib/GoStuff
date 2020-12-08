package main

import (
	"fmt"
	"strconv"
)

// variables can also be declared at package level & instead of declaring var everytime like this..
// var firstName string = "sakib"
// var lastName string = "alamin"
// var company string = "appscode"
// var id int = 12345

// .. we can enclose them in parenthesis like this.
// & ofcourse follow better naming convention, don't just put x, y etc as variable name
// use meaningful & considerable lenghth considering lifecycle of that particular variables
var (
	firstName string = "sakib"
	lastName  string = "alamin"
	company   string = "appscode"
	id        int    = 12345
)

// if we want to export something, then it has to follow a certain naming convention, must start with a Capital letter

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

	fmt.Println(firstName, lastName, company, id)

	// Main Data types -> notice that, there's no char data type in golang :(

	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	var isPrime bool = false
	fmt.Printf("%v, %T\n", isPrime, isPrime)

	const number int = 10
	// can't redefine the value as number is a constant like this, number += 10, come on, constant is constant!!

	fmt.Println(number + 10)

	// string to int converting, make sure that you import "strconv"
	s := "123"
	fmt.Printf("%T, %v\n", s, s)
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("some error occurred while conversion")
	} else {
		fmt.Println("no error, successfully converted!")
	}
	fmt.Printf("%T, %v\n",num, num)

	// int to string conversion, import "strconv" & look that Itoa returns only one value
	x := 123
	fmt.Printf("%T, %v\n", x, x)
	t := strconv.Itoa(x)
	fmt.Printf("%T, %v\n",t, t)
}
