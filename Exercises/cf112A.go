package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	for i := 0; i < len(a); i++ {
		x := strings.ToLower(a[i : i + 1])
		y := strings.ToLower(b[i : i + 1])
		if x == y {
			continue
		} else if x > y {
			fmt.Println(1)
			os.Exit(0)
		} else {
			fmt.Println(-1)
			os.Exit(0)
		}
	}

	fmt.Println(0)
}