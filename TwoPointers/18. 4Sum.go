package main

import (
	"slices"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	slices.Sort(nums)
	result, quad := [][]int{}, []int{}

	var kSum func(k, start, target int)
	kSum = func(k, start, target int) {
		if k != 2 {
			for i := start; i < len(nums)-k+1; i++ {
				if i > start && nums[i] == nums[i-1] {
					continue
				}

				quad = append(quad, nums[i])
				kSum(k-1, i+1, target-nums[i])
				quad = quad[:len(quad)-1]
			}
			return
		}
		l, r := start, len(nums)-1
		for l < r {
			s := nums[l] + nums[r]
			if s < target {
				l++
			} else if s > target {
				r--
			} else {
				result = append(result, append(quad, nums[l], nums[r]))
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	kSum(4, 0, target)

	return result
}

// func main() {
// 	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
// 	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))

// 	fmt.Println("Good")
// }
