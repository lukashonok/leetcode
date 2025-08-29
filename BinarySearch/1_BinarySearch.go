package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		mid := (left + right) / 2
		switch {
		case left > right:
			return -1
		case nums[mid] == target:
			return mid
		case target > nums[mid]:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
}
