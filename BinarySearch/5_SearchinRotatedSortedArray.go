package main

func searchInRoteted(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		mid := (left + right) / 2
		if left > right {
			return -1
		} else if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			if nums[left] > nums[mid] || target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] > nums[right] || target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
}
