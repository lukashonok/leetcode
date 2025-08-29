package main

func minCostClimbingStairs2(cost []int) int {
	m := make([]int, len(cost)+1)
	for i := len(m) - 1; i > 1; i-- {
		m[i-1] = m[i] + min(cost[i-1], cost[i-2])
	}
	return min(m[0], m[1])
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+2) // dp[n] и dp[n+1] будут 0 по умолчанию

	for i := n - 1; i >= 0; i-- {
		dp[i] = cost[i] + min(dp[i+1], dp[i+2])
	}
	return min(dp[0], dp[1])
}

// func main() {
// 	fmt.Println(minCostClimbingStairs([]int{1, 2, 3, 4, 5}))
// 	fmt.Println(minCostClimbingStairs2([]int{1, 2, 3, 4, 5}))
// 	// fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
// }
