package main

func findTargetSumWays(nums []int, target int) int {
	dp := map[int]map[int]int{}

	var dfs func(i int, currentSum int) int
	dfs = func(i, currentSum int) int {
		if i == len(nums) {
			if currentSum == target {
				return 1
			}
			return 0
		} else if _, ok := dp[i][currentSum]; ok {
			return dp[i][currentSum]
		}
		if _, ok := dp[i]; !ok {
			dp[i] = map[int]int{}
		}
		positive := dfs(i+1, currentSum+nums[i])
		negative := dfs(i+1, currentSum-nums[i])
		dp[i][currentSum] = positive + negative
		return dp[i][currentSum]
	}

	return dfs(0, 0)
}

// func main() {
// 	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
// }
