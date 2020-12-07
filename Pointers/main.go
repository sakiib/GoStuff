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

func setNewVal(val *int) {
	*val = 10
}

func main() {
	fmt.Println("Pointers!")
	currentVal := 10
	fmt.Println(currentVal)
	setValWithourReff(currentVal)
	fmt.Println("without pointer reference: ", currentVal)
	setValWithReff(&currentVal)
	fmt.Println("with pointer reference: ", currentVal)

	// using the new operator
	// new takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
	// In some programming languages there is a significant difference between using new and &, with great care being
	// needed to eventually delete anything created with new. Go is not like this, it's a garbage collected programming
	// language which means memory is cleaned up automatically when nothing refers to it anymore.
	newVal := new(int)
	fmt.Println(*newVal)
	setNewVal(newVal)
	fmt.Println(*newVal)
}
