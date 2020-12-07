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
		stack = stack[:topIndex]     // updating the stack, exluding the stack or simply we're popping the top of the stack
	}

	// declaring a queue of integer using make & slice
	queue := make([]int, 0)
	// appending or adding values to the queue
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 5)
	queue = append(queue, 10)
	queue = append(queue, 0)
	fmt.Println(queue, len(queue))

	// printing a queue
	for len(queue) > 0 {
		frontIndex := 0                // index of the value that is our current front of the queue which is 0
		fmt.Println(queue[frontIndex]) // printing the front of the queue
		queue = queue[frontIndex+1:]   // updating the queue, excluding the front or simply we're popping the front of the stack
	}
}
