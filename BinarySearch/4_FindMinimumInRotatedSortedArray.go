package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	mid := (right + left) / 2

	for left < right {
		if nums[left] > nums[right] {
			// There's rot
			if nums[mid] < nums[left] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			// Without rotation, smallest will be at left side anyway
			right = mid
		}

		mid = (right + left) / 2
	}

	return nums[left]
}
