package main

import (
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	fmt.Println("Concurrency!")
	// concurrency != parallelism
	// Concurrency is when two or more tasks can start, run, and complete in overlapping time periods.
	// It doesn't necessarily mean they'll ever both be running at the same instant. For example, multitasking
	// on a single-core machine. Parallelism is when tasks literally run at the same time, e.g., on a multicore processor.

	// Goroutines
	f(0)
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
