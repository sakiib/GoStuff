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

func func2() int {
	return 100
}

func func1() int {
	return func2()
}

func multiValRetFunc() (int, int) {
	return 5, 10
}

func main() {
	fmt.Println("Functions!")
	fmt.Println(sumOfTwoNumbers(10, 20))
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(sumOfAList(slice1))
	// functions are built-up in a stack like fashion
	fmt.Println(func1())
	// a function can return multiple values, types must given in similar fashion
	val1, val2 := multiValRetFunc()
	fmt.Println(val1, val2)
}
