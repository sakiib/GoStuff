package main

import (
	"fmt"
)

func checkBracketSquence(s string) string {
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		cur := s[i : i+1]
		if cur == "(" {
			stack = append(stack, cur)
		} else {
			if len(stack) > 0 {
				topIndex := len(stack) - 1
				if stack[topIndex] == "(" {
					stack = stack[:topIndex]
				} else {
					stack = append(stack, ")")
				}
			} else {
				stack = append(stack, ")")
			}
		}
	}
	if len(stack) == 0 {
		return "balanced!"
	} else {
		return "not balanced!"
	}
}

func main() {
	fmt.Println("balanced parenthesis!")
	s := "((()))"
	fmt.Println("this given sequence is: ", checkBracketSquence(s))
}
