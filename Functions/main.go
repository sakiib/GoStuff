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

// variadic function, can take any number of arguments as the last parameter only
// By using ... before the type name of the last parameter we can indicate that it takes zero or more of those parameters.
// This is precisely how the fmt.Println function is implemented

func sumUsingVariadicFunc(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}

// no built-int function for calculating minimum, variadic function can be used here to implement one
// it takes two or more arguments & calculates the minimum value among them, max can be implemented similarly

func min(x int, args ...int) int {
	mn := x
	for _, value := range args {
		if value < mn {
			mn = value
		}
	}
	return mn
}

func max(x int, args ...int) int {
	mx := x
	for _, value := range args {
		if value > mx {
			mx = value
		}
	}
	return mx
}

// a simple recursive function to calculate the factorial of a number
func recFunc(n int) int {
	if n < 0 {
		panic("negative number in factorial function")
	}
	if n <= 1 {
		return 1
	}
	return n * recFunc(n-1)
}

// Higher order function, see the def. in the main function
func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
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

	fmt.Println(sumUsingVariadicFunc(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(min(5, 4, 2, 1, 3), max(1, 4, 5, 2, 3))

	fmt.Println(recFunc(5), recFunc(0))

	// Closure -> it is possible to create functions insides of functions, like this:
	// here, add is a local variable that has the type func(int, int) int (a function that takes two ints and returns an int).
	addTwo := func(x, y int) int {
		return x + y
	}
	fmt.Println(addTwo(10, 20))
	// When you create a local function like this it also has access to other local variables, should print 1, 2
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment(), increment())

	swap := func(x, y int) (int, int) {
		return y, x
	}
	// swap two integers
	x, y := swap(10, 20)
	fmt.Println(x, y)
	// or we can just do this! cute! :p
	a := 100
	b := 200
	a, b = b, a
	fmt.Println(a, b)

	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	newInts := intSeq()    // redifining i = 0
	fmt.Println(newInts()) // 1

	// Higher order functions: The definition of Higher-order function:
	// a function which does at least one of the following: i. takes one or more functions as arguments ii. returns a function as its result
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}
