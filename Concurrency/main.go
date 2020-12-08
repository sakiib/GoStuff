package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * 3000)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	fmt.Println("Concurrency!")
	// concurrency != parallelism
	// Concurrency is when two or more tasks can start, run, and complete in overlapping time periods.
	// It doesn't necessarily mean they'll ever both be running at the same instant. For example, multitasking
	// on a single-core machine. Parallelism is when tasks literally run at the same time, e.g., on a multicore processor.

	// Goroutines
	go f(0)
	//var input string
	//fmt.Scanln(&input)
	//fmt.Println("input val: ", input)

	// Channels: Channels provide a way for two goroutines to communicate with one another and synchronize their execution.
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
