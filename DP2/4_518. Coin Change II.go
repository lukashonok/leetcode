package main

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][amount] = 1
	}

	for j := amount - 1; j >= 0; j-- {
		for i := len(coins) - 1; i >= 0; i-- {
			dp[i][j] = dp[i+1][j]
			if amount-j-coins[i] >= 0 {
				dp[i][j] += dp[i][j+coins[i]]
			}
		}
	}

	return dp[0][0]
}

// func main() {
// 	fmt.Println(change(5, []int{1, 2, 5}))
// }
