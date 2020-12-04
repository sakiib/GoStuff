package main

import (
	"fmt"
)

// basic struct: func functionName(argument argument type, ) return type {}
func sumOfTwoNumbers(x, y int) int {
	return x + y
}

func sumOfAList(tempList []int) int {
	sum := 0
	for _, val := range tempList {
		fmt.Println("here: ", val)
		sum += val
	}
	return sum
}

func main() {
	fmt.Println("Functions!")
	fmt.Println(sumOfTwoNumbers(10, 20))
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(sumOfAList(slice1))
}
