package main

import (
	"fmt"
)

func main() {
	fmt.Println("Data Structures!")

	// delcaring a stack of integer using make & slice
	stack := make([]int, 0)
	// appending or adding values to the stack
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 5)
	stack = append(stack, 10)
	stack = append(stack, 0)
	fmt.Println(stack, len(stack))

	// printing a stack
	for len(stack) > 0 {
		topIndex := len(stack) - 1   // index of the value that is our current top of the stack
		fmt.Println(stack[topIndex]) // printing the top of the stack
		stack = stack[:topIndex]     // updating the stacking, exluding the stack or simply we're popping the top of the stack
	}

}
