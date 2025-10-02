package main

func canPartition(nums []int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total%2 == 1 {
		return false
	}
	target := total / 2
	dp := map[int]struct{}{0: {}}
	for i := len(nums) - 1; i >= 0; i-- {
		dpClone := map[int]struct{}{}
		for s := range dp {
			dpClone[s+nums[i]] = struct{}{}
			dpClone[s] = struct{}{}
			_, exists := dpClone[target]
			if exists {
				return true
			}
		}
		dp = dpClone
	}
	_, exists := dp[target]
	return exists
}

func change(amount int, coins []int) int {
	dp := map[int]int{}

	var dfs func(price int) int
	dfs = func(price int) int {
		if _, ok := dp[price]; ok {
			return dp[price]
		}
		total := 0
		for _, coin := range coins {
			if price-coin == 0 {
				total++
			} else if price-coin > 0 {
				total += dfs(price - coin)
			}
		}
		dp[price] = total
		return total
	}

	return dfs(amount)
}

// func main() {
// 	fmt.Println(canPartition([]int{1, 5, 11, 5}))
// }
