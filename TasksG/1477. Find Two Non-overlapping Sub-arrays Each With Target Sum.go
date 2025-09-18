package main

import (
	"math"
)

func minSumOfLengths(arr []int, target int) int {
	if len(arr) < 2 {
		return -1
	}
	dp := make([]int, len(arr)+1)
	for i := range dp {
		dp[i] = math.MaxInt32 / 2
	}

	windowSum, optimal, curOptimal := 0, math.MaxInt32/2, math.MaxInt32/2
	l, r := 0, 0
	for r < len(arr) {
		windowSum += arr[r]
		r++

		if windowSum > target {
			for windowSum > target {
				windowSum -= arr[l]
				l++
			}
		}

		if windowSum == target {
			size := r - l
			optimal = min(optimal, size+dp[l])
			curOptimal = min(curOptimal, size)
		}

		dp[r] = curOptimal
	}

	if optimal == math.MaxInt32/2 {
		return -1
	}

	return optimal
}

// func main() {
// 	fmt.Println(minSumOfLengths([]int{1, 6, 1}, 7))
// 	fmt.Println(minSumOfLengths([]int{4, 3, 2, 6, 2, 3, 4}, 6))
// 	fmt.Println(minSumOfLengths([]int{7, 3, 4, 7}, 7))
// 	fmt.Println(minSumOfLengths([]int{3, 1, 1, 1, 5, 1, 2, 9}, 3))
// 	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3))
// }
