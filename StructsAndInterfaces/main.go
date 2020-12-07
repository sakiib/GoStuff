package main

import (
	"fmt"
)

// Student struct
type Student struct {
	ID        string
	FirstName string
	LastName  string
	CG        float32
	Credits   int
}

func main() {
	fmt.Println("structs and interfaces!")
	student1 := Student{ID: "123", FirstName: "sakib", LastName: "alamin", CG: 3.5, Credits: 130}
	// %v prints only the values, if we need the name of the fields, we can use %+v to achieve that
	fmt.Printf("%+v\n", student1)
	fmt.Println(student1)
}
