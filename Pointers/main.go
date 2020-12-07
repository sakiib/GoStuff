package main

import (
	"fmt"
)

func setValWithourReff(val int) {
	val = 0
}

func setValWithReff(val *int) {
	*val = 0
}

func main() {
	fmt.Println("Pointers!")
	currentVal := 10
	fmt.Println(currentVal)
	setValWithourReff(currentVal)
	fmt.Println("without pointer reference: ", currentVal)
	setValWithReff(&currentVal)
	fmt.Println("with pointer reference: ", currentVal)
}
