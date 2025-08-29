package main

import (
	"slices"
)

func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(dp) <= 2 {
		return slices.Max(nums)
	} else {
		dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	}
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return max(dp[len(dp)-1], dp[len(dp)-2])
}

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

// func main() {
// 	fmt.Println(rob([]int{1, 4, 1, 1, 5}))
// 	fmt.Println(rob([]int{2, 1, 1, 2}))
// }
