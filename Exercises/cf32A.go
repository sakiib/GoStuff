package main

import (
	"fmt"
)

func abs(x int, y int) int {
	if x >= y {
		return x - y
	}  else {
		return y - x
	}
}

func main() {
	var (
		n int
		d int
	)
	fmt.Scan(&n, &d)

	var ara []int
	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val)

		ara = append(ara, val)
	}

	var ans int = 0
	for i, x := range ara {
		for j, y := range ara {
			if i != j && abs(x, y) <= d {
				ans ++
			}
		}
	}
	fmt.Println(ans)
}