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
	Courses   []string
}

func main() {
	fmt.Println("structs and interfaces!")

	// create a list of courses & then assign or we can assign in place like line 23
	sub := []string{"maths", "cse"}
	fmt.Println(sub)
	//student1 := Student{ID: "123", FirstName: "sakib", LastName: "alamin", CG: 3.5, Credits: 130, Courses: sub}
	student1 := Student{ID: "123", FirstName: "sakib", LastName: "alamin", CG: 3.5, Credits: 130, Courses: []string{"maths", "cse"}}

	// %v prints only the values, if we need the name of the fields, we can use %+v to achieve that
	fmt.Printf("%+v\n", student1)
	fmt.Println(student1)

	// declaring an empty slice of Student type, notice the opening & closing curly braces at the end
	studentsList := []Student{}
	studentsList = append(studentsList, Student{ID: "123", FirstName: "sakib", LastName: "alamin", CG: 3.5, Credits: 130, Courses: []string{"maths", "cse"}})
	studentsList = append(studentsList, Student{ID: "234", FirstName: "prangan", LastName: "majumder", CG: 3.6, Credits: 120, Courses: []string{"english", "cse"}})
	fmt.Println(studentsList)

	// we can access the field values using the dot(.) operator, courses is a list or slice, so print like this.
	for index, student := range studentsList {
		fmt.Println(index, student.FirstName, student.LastName, student.CG)
		for _, course := range student.Courses {
			fmt.Print(course, " ")
		}
		fmt.Println()
	}

	// var studentsArray []Student
	// studentsArray = []Student or the one below. an array of Student Items
	studentsArray := []Student{
		Student{ID: "123", FirstName: "sakib", LastName: "alamin", CG: 3.5, Credits: 130, Courses: []string{"maths", "cse"}},
		Student{ID: "234", FirstName: "prangan", LastName: "majumder", CG: 3.6, Credits: 120, Courses: []string{"english", "cse"}},
	}

	fmt.Println(studentsArray)
}
