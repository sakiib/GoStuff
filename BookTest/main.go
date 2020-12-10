package main

import (
	"fmt"
	"sort"
)

type Book struct {
	ID int
	Title string
	Author string
	Genre string
	Rating int
}

var books []Book

// function to print the number of books available & print the books details
func getBooks() {
	fmt.Println("func: getBooks")
	fmt.Println("Total books available right now: ", len(books))
	for _, book := range books {
		fmt.Printf("%+v\n", book)
	}
}

// function to get a single book by it's ID
func getBook(id int) Book {
	fmt.Println("func: getBook")
	for _, book := range books {
		if book.ID == id {
			return  book
		}
	}
	return Book{}
}

// add a single book to the list of books
func addBook(book Book) {
	fmt.Println("func: addBook")
	books = append(books, book)
}

// update a books with the same ID with the given book & details
func updateBook(newBook Book) {
	fmt.Println("func: updateBook")
	for index, book := range books {
		if book.ID == newBook.ID {
			books = append(books[:index], books[index + 1:]...)
			books = append(books, newBook)
			return
		}
	}
}

// delete a book with the given ID
func deleteBook(id int) {
	fmt.Println("func: getBook")
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index + 1:]...)
			return
		}
	}
}

func main() {
	fmt.Println("Book Test Operations!")

	books = []Book{
		Book{ID: 1, Title: "book - 1", Author: "sakib", Genre: "comedy", Rating: 4},
		Book{ID: 3, Title: "book - 3", Author: "prangon", Genre: "romcom", Rating: 5},
		Book{ID: 4, Title: "book - 4", Author: "mehedi", Genre: "action", Rating: 3},
		Book{ID: 2, Title: "book - 2", Author: "sahadat", Genre: "action", Rating: 3},
	}

	//books = append(books, Book{ID: 1, Title: "book - 1", Author: "sakib", Genre: "comedy", Rating: 4})
	//books = append(books, Book{ID: 3, Title: "book - 3", Author: "prangon", Genre: "romcom", Rating: 5})
	//books = append(books, Book{ID: 4, Title: "book - 4", Author: "mehedi", Genre: "action", Rating: 3})
	//books = append(books, Book{ID: 2, Title: "book - 2", Author: "sahadat", Genre: "action", Rating: 3})

	//for _, book := range books {
	//	fmt.Printf("%+v\n", book)
	//}

	// sorting the books[] with Rating in descending & then with ID in ascending order
	sort.SliceStable(books, func(i, j int) bool {
		if books[i].Rating == books[j].Rating {
			return books[i].ID < books[j].ID
		} else {
			return books[i].Rating > books[j].Rating
		}
	})

	//for _, book := range books {
	//	fmt.Printf("%+v\n", book)
	//}

	getBooks()
	book := getBook(10)
	fmt.Printf("%+v\n", book)
	addBook(Book{ID: 5, Title: "book - 5", Author: "TS", Genre: "thriller", Rating: 5})
	deleteBook(2)
	getBooks()
	deleteBook(10)
	getBooks()
	updateBook(Book{ID: 5, Title: "book - 5 - five", Author: "Tamal Saha", Genre: "thriller & mystery", Rating: 5})
	getBooks()
}