package main

import (
	"fmt"
	"strconv"
)

func main() {
	var testCase int
	fmt.Scan(&testCase)
	for tc := 1; tc <= testCase; tc++ {
		var word string
		fmt.Scan(&word)
		if len(word) > 10 {
			firstChar := word[0 : 1]
			lastChar := word[len(word) - 1 : len(word)]
			fmt.Println(firstChar + strconv.Itoa(len(word) - 2) + lastChar)
		} else {
			fmt.Println(word)
		}
	}
}