package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)
	str += "+"

	var numbers []int
	temp := 0
	for i := 0; i < len(str); i++ {
		cur := str[i : i+1]
		if cur == "+" {
			numbers = append(numbers, temp)
			temp = 0
		} else {
			x, err := strconv.Atoi(cur)
			if err != nil {
				fmt.Println("error")
				os.Exit(0)
			}
			temp = temp*10 + x
		}
	}

	sort.Ints(numbers)
	for idx, val := range numbers {
		if idx == 0 {
			fmt.Print(val)
		} else {
			fmt.Print("+")
			fmt.Print(val)
		}
	}

	fmt.Println()
}