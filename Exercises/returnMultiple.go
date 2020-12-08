package main

import (
	"fmt"
)

func returnMul(n int) (int, bool) {
	reminder := n%2
	isEven := false
	if reminder == 0 {
		isEven = true
	}
	return reminder, isEven
}
func main() {
	for i := 1; i <= 10; i++ {
		reminder, isEven := returnMul(i)
		fmt.Printf("%+v %+v\n", reminder, isEven)
	}
}
