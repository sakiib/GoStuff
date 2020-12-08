package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		cur := s[i : i+1]
		if cur == "." {
			fmt.Print(0)
		} else {
			i++
			nxt := s[i : i+1]
			if nxt == "." {
				fmt.Print("1")
			} else {
				fmt.Print("2")
			}
		}
	}
}