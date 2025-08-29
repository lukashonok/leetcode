package main

func lengthOfLIS(nums []int) int {
	lis, result := make([]int, len(nums)), 1
	lis[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		lis[i] = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				lis[i] = max(lis[j]+1, lis[i])
			}
		}
		result = max(result, lis[i])
	}
	return result
}

// func main() {
// 	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
// 	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
// 	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
// }
