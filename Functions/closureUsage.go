package main

import (
	"fmt"
)

// 1. Isolating data
func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	fmt.Println("makeFibGen func assigned as closure!")
	return func() int {
		fmt.Println("inside the anonymous func!")
		f2, f1 = (f1 + f2), f2
		return f1
	}
}

func main() {
	fmt.Println("closure usage!")
	// 1. Isolating data
	gen := makeFibGen()
	fmt.Println("we have the gen func!")
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
}
