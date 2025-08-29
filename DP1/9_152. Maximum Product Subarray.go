package main

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return max(0, nums[0])
	}
	result := nums[0]
	dp := make([][]int, len(nums))
	dp[0] = []int{nums[0], nums[0]}
	for i := 1; i < len(nums); i++ {
		minProd, maxProd := nums[i]*dp[i-1][0], nums[i]*dp[i-1][1]
		minProd, maxProd = min(nums[i], minProd, maxProd), max(nums[i], minProd, maxProd)
		dp[i] = []int{minProd, maxProd}
		result = max(result, maxProd)
	}
	return result
}

// func main() {
// 	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
// }
