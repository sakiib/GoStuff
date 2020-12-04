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

	// length & capacity -> they aren't same
	fmt.Println(len(a), cap(a))

	aa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bb := aa[:]   // slice of all elements
	cc := aa[3:]  // slice from 4th element to end
	dd := aa[:6]  // slice first 6 elements, so basically [left, right), left is inclusive & right is exclusive
	ee := aa[3:6] // slice the 4th, 5th, and 6th elements, aa, bb, cc, dd, ee all point to same underlying data
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println(ee)

	// 2D slices, how we store graph. e.g: vector <int> graph[N], each of the indices from [0, 10) have an empty slice with them
	var graph [10][]int
	for i := 0; i < 10; i++ {
		fmt.Println(graph[i])
	}

	// appending at the end of a slice, we can simply append a single value, but to insert a slide we need to use the (...) spread operator/syntax
	var testSlice []int
	testSlice = append(testSlice, 10)
	testSlice = append(testSlice, 20)
	fmt.Println("testSlice: ", testSlice)
	var addSlice []int
	addSlice = append(addSlice, 30)
	addSlice = append(addSlice, 40)
	fmt.Println("addSlice", addSlice)
	testSlice = append(testSlice, addSlice...) // see the three dot's (spread operator/syntax)
	fmt.Println("updated testSlice", testSlice)
	fmt.Println("see if the addSlice has changed: ", addSlice) // no, it hasn't
	// declaration of a map [key, value]: [string, int] using the make function
	mp := make(map[string]int)
	mp["sakib"] = 1
	mp["alamin"] = 2
	mp["someone"] = 3
	// printing the map using key, val pair & ranging over the map
	for key, val := range mp {
		fmt.Println(key, val)
	}
	fmt.Println(mp["non_existant_key"]) // so, non existent key gives zero, so how do we check if it's actually there or not!!!

	// so, to ensure that the key exists in our map, ok must be equal to true
	val, ok := mp["someone"]
	if ok {
		fmt.Println("key exists & it's value is", val)
	} else {
		fmt.Println("key doesn't exist in the map mp")
	}
	// this is cute!! so, here's a semicolon!! we can use this condition to check the existance & print the value if it's present
	// remember, them scope of this value is inside the curly braces only
	if value, ok := mp["someone"]; ok {
		fmt.Println("val: ", value)
	} else {
		fmt.Println("...not present")
	}
	// for some reason if we don't actually need the value & we just only need to check the existance then we can use underscore(_)
	if _, ok := mp["non_existant_key"]; ok {
		fmt.Println("the key exists")
	} else {
		fmt.Println("...not present")
	}
}
