package main

import (
	"fmt"
)

func main() {
	fmt.Println("go nuances!")
	mySlice := make([]int, 5)
	fmt.Println(mySlice)
	mySlice = append(mySlice, 10)
	fmt.Println(mySlice)

	// Slice Copy: copying mySlice to cpySlice, remember it has a reference to the mySlice, so will get modified along with mySlice
	cpySlice := make([]int, len(mySlice))
	cpySlice = mySlice
	fmt.Println(cpySlice)

	// Slice Cut: new cutSlice should be [1, 2, 3, 8, 9, 10], see the ... spread operator at the end as we can't just append a slice at the end
	cutSlice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cutSlice = append(cutSlice[:3], cutSlice[7:]...)
	fmt.Println(cutSlice)
}