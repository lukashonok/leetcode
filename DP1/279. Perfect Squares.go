package main

import (
	"fmt"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = n
	}
	dp[0] = 0

	for target := 1; target <= n; target++ {
		for s := 1; s <= target; s++ {
			square := s * s
			if target-square < 0 {
				break
			}
			dp[target] = min(dp[target], 1+dp[target-square])

		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquares(12)) // 3 (4, 4, 4)
	fmt.Println(numSquares(13)) // 2 (4, 9)
}
