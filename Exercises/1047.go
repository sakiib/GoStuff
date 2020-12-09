package main

import (
	"fmt"
)

const INF int = 2e9

func minFunc(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var col [25][25] int
	var dp[25][25] int

	// todo: input section

	n := 20

	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			dp[i][j] = INF
		}
	}

	for i := 1; i <= 3; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				if j != k {
					dp[i][j] = minFunc(dp[i][j], col[i][j] + dp[i - 1][k])
				}
			}
		}
	}

	ans := INF
	for i := 1; i <= 3; i++ {
		ans = minFunc(ans, dp[n][i])
	}

	fmt.Println(ans)
}
