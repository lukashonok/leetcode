package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	left, right := 1, len(matrix)

	for left < right {
		mid := (right + left) / 2
		if matrix[mid][0] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return search(matrix[left-1], target) >= 0
}
