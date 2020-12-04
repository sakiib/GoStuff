package main

import (
	"fmt"
)

// basic struct: func functionName(argument argument type, ) return type {}
func sumOfTwoNumbers(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Functions!")
	fmt.Println(sumOfTwoNumbers(10, 20))
}
