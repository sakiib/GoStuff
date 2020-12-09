package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("file IO handling!")

	// Checking if a file exists or not: Stat returns file info., returns an error if there is no file.
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does exist. File information:", fileInfo.Name())

	// reading file using ioutil lib.
	fileContent, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", fileContent) // fileContent -> all the characters from the file
	for _, val := range fileContent {
		if val == '\n' {
			fmt.Println()
		} else {
			fmt.Print(string(val))
		}
	}
}