package main

import (
	"fmt"
)

const inf int = 2e9

var dp = make(map[int]int)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getDivisors(n int) []int {
	var divisors []int
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if n/i != i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}

func solve(n int) int {
	if n == 1 {
		return 0
	}
	if _, ok := dp[n]; ok {
		return dp[n]
	}
	ret := inf
	divisors := getDivisors(n)
	ret = min(ret, solve(n-1)+1)
	for _, d := range divisors {
		ret = min(ret, solve(n/d)+1)
	}
	dp[n] = ret
	return dp[n]

}
func main() {
	var t int
	fmt.Scan(&t)

	for tc := 1; tc <= t; tc++ {
		var n int
		fmt.Scan(&n)
		// dp = nil
		// fmt.Println("TLE: ", solve(n))
		if n <= 3 {
			fmt.Println(n - 1)
		} else {
			if (n-3)%2 == 0 {
				fmt.Println(3)
			} else {
				fmt.Println(2)
			}
		}
	}
}
