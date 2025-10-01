package main

import "fmt"

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		left, mid, right := nums[l], nums[m], nums[r]

		if left == target || mid == target || right == target {
			return true
		}

		if left < mid {
			// we're in left portion
			if left <= target && mid > target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if left > mid {
			// we're in right portion
			if mid < target && right >= target {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			l++
			// we can'd determine where we are, so just shift left index
		}
	}
	return false
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 1))
}
