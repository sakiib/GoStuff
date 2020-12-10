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

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
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
	var input1 string
	fmt.Scanln(&input1)
	fmt.Println("input val: ", input1)

	// Channels: Channels provide a way for two goroutines to communicate with one another and synchronize their execution.
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input2 string
	fmt.Scanln(&input2)

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
	//Explanation:
	//numbers goroutine time stamps: 0 250 500 750 1000 1250 (ms)
	//alphabets goroutine time stams: 0 400 800 1200 1600 2000 (ms)
	//main goroutine: 3000 ms
	//overall: 0 250 400 500 750 800 1000 1200 1250 1600 2000 3000 (ms)
	//output: 1 a 2 3 b 4 c 5 d e main terminated
}
