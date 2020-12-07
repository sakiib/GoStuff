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

// Book struct
type Book struct {
	ID     string
	Title  string
	Isbn   string
	Author *Author
}

// Author struct
type Author struct {
	ID        string
	FirstName string
	LastName  string
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

	// other ways of declaration
	var st Student
	fmt.Println(st.FirstName, st.LastName, st.Courses, st.CG, st.Credits)

	// using thie new()
	// This allocates memory for all the fields, sets each of them to their zero value and returns a pointer.
	// (*Student) More often we want to give each of the fields a value. We can do this in two ways. Like this: stdnt := Student{fieldName: initialValue}
	stdnt := new(Student)
	fmt.Println(stdnt.FirstName, stdnt.LastName, stdnt.Courses, stdnt.CG, stdnt.Credits)

	// Book - data, author refernce to another struct Author
	book1 := Book{ID: "1", Title: "book-1", Isbn: "1354", Author: &Author{ID: "12", FirstName: "kazi", LastName: "Nazrul"}}
	fmt.Println(book1.ID, book1.Title, book1.Isbn, book1.Author.FirstName, book1.Author.LastName, book1.Author.ID)

	booksArray := []Book{
		Book{ID: "123", Title: "book-1", Isbn: "12345", Author: &Author{ID: "12", FirstName: "kazi", LastName: "Nazrul"}},
		Book{ID: "234", Title: "book-2", Isbn: "56789", Author: &Author{ID: "34", FirstName: "rabindranath", LastName: "tagore"}},
	}
	fmt.Println(booksArray)
}
