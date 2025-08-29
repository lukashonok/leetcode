package main

func maxSubArray(nums []int) int {
	maxSum, curSum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		curSum += nums[i]
		curSum = max(nums[i], curSum)
		maxSum = max(curSum, maxSum)
	}
	return maxSum
}

// func main() {
// 	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
// 	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // 23
// 	fmt.Println(maxSubArray([]int{1}))                             // 1
// 	fmt.Println(maxSubArray([]int{-1}))                            // -1
// }
