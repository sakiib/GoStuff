package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	// according to the book..
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// use defer file.Close() right after opening the file to make sure the file is closed as soon as the function completes.
	// see the definitions of defer if needed

		// get the file size
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bs)
	fmt.Println(str)

	// using the ioutil
	fs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	st := string(fs)
	fmt.Println(st)

	// Creating a file
	cfile, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer cfile.Close()
	file.WriteString("test")

	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	// To get the contents of a directory we use the same os.
	//Open function but give it a directory path instead of a file name. Then we call the Readdir method
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// Often we want to recursively walk a folder (read the folder's contents, all the sub-folders, all the
	// sub-sub-folders, …). To make this easier there's a Walk function provided in the path/filepath package
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}