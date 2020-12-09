package main

import (
	"fmt"
)

func main() {
	fmt.Println("Control Structures!")

	// for loop & only type of loop in golang! cute! :p
	for i := 1; i <= 10; i++ {
		fmt.Println("I'm a golang for loop: ", i)
	}

	// this is another way to write the for loop
	idx := 1
	for idx <= 10 {
		fmt.Println("I'm also a for loop: ", idx)
		idx++
	}

	// conditionals: if - else if - else
	for i := 1; i <= 10; i++ {
		if i%6 == 0 {
			fmt.Println(i, " is divisible by both 2 & 3")
		} else if i%2 == 0 {
			fmt.Println(i, " is divisible by both 2")
		} else if i%3 == 0 {
			fmt.Println(i, " is divisible by 3")
		} else {
			fmt.Println(i, " is not divisible by 2 or 3")
		}
	}

	// switch case:
	for i := 1; i <= 10; i++ {
		reminder := i % 2
		switch reminder {
		case 0:
			fmt.Println("Even")
		case 1:
			fmt.Println("Odd")
		}
	}

	// n Go language, rune literal expressed as one or more characters enclosed in single quotes like ‘g’, ‘\t’, etc.
	name := "Bangladesh"
	for index, v := range []rune(name) {
		fmt.Printf("%v %c %T\n", index, v, v)
	}
}
