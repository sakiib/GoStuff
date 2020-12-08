package main

import (
	"fmt"
)

var fibCalculated = make(map[int]int)

func nthFibonacci(n int) int {
	if n < 0 {
		panic("WTF!")
	}
	if n <= 1 {
		return n
	}
	if val, ok := fibCalculated[n]; ok {
		return val
	}
	fibCalculated[n] = nthFibonacci(n - 1) + nthFibonacci(n - 2)
	return fibCalculated[n]
}
func main() {
	fmt.Println(fibCalculated)
	for i := 1; i <= 50; i++ {
		fmt.Println(nthFibonacci(i))
	}
}