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

	// assigning values while declaring the array, filled from index 0, rests are filled with default value, which is also zero
	tempAra := [10]int{98, 93, 77, 82, 83}
	fmt.Println(tempAra)
	// we can also break them down in separate lines, in that case an ending comma is a must
	words := [5]string{
		"gone",
		"with",
		"the",
		"wind",
		"again",
	}
	fmt.Println(words)
	// we can declare an array regarding of it's size by using three dots like the code below
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers, " & lenghth is ", len(numbers))

	// copying an array is just like assigning a value to a variable, doesn't point to the originl array while assigning, it just copies it!
	aarray := [...]int{1, 2, 3}
	carray := aarray
	carray[1] = 10
	// although we've changed the content on carray, the value on aarray remains the same as before
	fmt.Println(aarray) // prints: [1, 2, 3]
	fmt.Println(carray) // prints: [1, 10, 3]

	// 2-D array declaration, basically created r by c empty arrays, now initiating 0 to r - 1 cells to assign new values
	var matrix [3][3]int
	matrix[0] = [3]int{1, 0, 0}
	matrix[1] = [3]int{0, 1, 0}
	matrix[2] = [3]int{0, 0, 1}
	fmt.Println(matrix)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	// for slices the size need not to be defined, it's more like C++ container vector <> in terms of usage but only better
	// unlike array, slices actually points to the array a, from which b is assigned, so both go changed, like pointers do
	aslices := []int{1, 2, 3}
	bslices := aslices
	bslices[1] = 10
	fmt.Println(aslices) // 1, 10, 3
	fmt.Println(bslices) // 1, 10, 3

	// if you want to create a slice you should use the builtin function "make" like this:
	mySlice := make([]int, 10)
	fmt.Println(mySlice)
	// This creates a slice that is associated with an underlying float64 array of length 5.
	// Slices are always associated with some array, and although they can never be longer than the array, they can be smaller.
	// The make function also allows a 3rd parameter, like this:
	newSlice := make([]int, 10, 20)
	// Here, 20 represents the capacity of the underlying array which the slice points to:
	fmt.Println(newSlice)
	// Clearing a slice
	a := []string{"A", "B", "C", "D", "E"}
	a = nil
	fmt.Println(a, len(a), cap(a)) // [] 0 0
	// In practice, nil slices and empty slices can often be treated in the same way: they have zero length and capacity,
	// they can be used with the same effect in for loops and append functions, and they even look the same when printed.
}
