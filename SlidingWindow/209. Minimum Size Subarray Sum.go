package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	minimal, l, r, currentSum, length := math.MaxInt, 0, 0, nums[0], len(nums)
	for l <= r {
		if currentSum < target && r < length-1 {
			r++
			currentSum += nums[r]
			continue
		} else if currentSum >= target {
			minimal = min(minimal, r-l+1)
		}
		currentSum -= nums[l]
		l++
	}
	if minimal == math.MaxInt {
		return 0
	}
	return minimal
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(7, []int{2}))
	fmt.Println(minSubArrayLen(2, []int{2}))
}
