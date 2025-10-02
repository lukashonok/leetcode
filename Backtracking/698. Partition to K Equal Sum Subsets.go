package main

import (
	"fmt"
)

func canPartitionKSubsets(nums []int, k int) bool {
	total, used := 0, make([]bool, len(nums))
	for _, v := range nums {
		total += v
	}
	target := total / k
	if total%k != 0 {
		return false
	}

	var dfs func(i, k, subsetSum int) bool
	dfs = func(i, k, subsetSum int) bool {
		if k == 0 {
			return true
		}
		if subsetSum == target {
			return dfs(0, k-1, 0)
		}

		for j := i; j < len(nums); j++ {
			if used[j] || nums[j]+subsetSum > target {
				continue
			}
			used[j] = true
			if dfs(j+1, k, nums[j]+subsetSum) {
				return true
			}
			used[j] = false
		}

		return false
	}

	return dfs(0, k, 0)
}

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
}
