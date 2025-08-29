package main

func maxProfit(prices []int) int {
	dp, length := map[int]map[bool]int{}, len(prices)

	var dfs func(i int, buying bool) int
	dfs = func(i int, buying bool) int {
		if i >= length {
			return 0
		}
		if _, ok := dp[i][buying]; ok {
			return dp[i][buying]
		}
		if _, ok := dp[i]; !ok {
			dp[i] = map[bool]int{}
		}
		cooldown := dfs(i+1, buying)
		if buying {
			buy := dfs(i+1, !buying) - prices[i]
			dp[i][buying] = max(cooldown, buy)
		} else {
			sell := dfs(i+2, !buying) + prices[i]
			dp[i][buying] = max(cooldown, sell)
		}
		return dp[i][buying]
	}

	return dfs(0, true)
}

// func main() {
// 	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
// }
