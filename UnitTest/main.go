package main

import "fmt"

func main() {
	fmt.Println("Unit Test!")
	fmt.Println(Total(1, 3, 5, 6))
	fmt.Println(Min(20, 500))
	fmt.Println(Max(10, 30))
}

func Total(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}