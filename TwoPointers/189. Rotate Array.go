package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}

	count := 0
	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % n
			nums[next], prev = prev, nums[next]
			current = next
			count++
			if current == start {
				break
			}
		}
	}
	fmt.Println(nums)
}

// func main() {
// 	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 10)
// 	rotate([]int{-1, -100, 3, 99}, 2)
// }
