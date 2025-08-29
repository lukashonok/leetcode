package main

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for {
		mid := (left + right) / 2

		leftNeighbor, rightNeighbor := mid-1, mid+1

		if leftNeighbor < 0 {
			leftNeighbor = 0
		}
		if rightNeighbor == len(nums) {
			rightNeighbor = len(nums) - 1
		}
		if nums[leftNeighbor] <= nums[mid] && nums[rightNeighbor] <= nums[mid] {
			return mid
		}
		if nums[mid] < nums[leftNeighbor] {
			right = mid - 1
		} else if nums[mid] < nums[rightNeighbor] {
			left = mid + 1
		}
	}
}

// func main() {
// 	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))          // 2
// 	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})) // 5
// 	fmt.Println(findPeakElement([]int{3, 1, 2}))             // 0
// }
