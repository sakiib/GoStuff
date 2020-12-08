package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	m := make(map[string] int)
	r := bufio.NewReader(os.Stdin)
	var test int
	fmt.Fscanf(r, "%d\n", &test)
	for testCase := 1; testCase <= test; testCase++ {
		var str string
		fmt.Fscanf(r, "%s\n", &str)
		val, ok := m[str]
		if !ok {
			fmt.Printf("OK\n")
		} else {
			fmt.Printf("%s%d\n", str, val)
		}
		m[str] ++
	}
}

// using the bufio.NewReader(os.Stdin)