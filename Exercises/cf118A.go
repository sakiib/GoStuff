package main

import (
	"fmt"
	"strings"
)

func isVowel(ch string) bool {
	if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" || ch == "y" {
		return true
	} else {
		return false
	}
}

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		currentChar := s[i : i + 1]
		if isVowel(strings.ToLower(currentChar)) {
			continue
		} else {
			fmt.Print("." + strings.ToLower(currentChar))
		}
	}

}