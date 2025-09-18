package main

func maxSubarraySumCircular(nums []int) int {
	minGlobal, maxGlobal := nums[0], nums[0]
	minCur, maxCur := 0, 0
	total := 0
	for _, n := range nums {
		total += n
		minCur, maxCur = min(n, n+minCur), max(n, n+maxCur)
		minGlobal, maxGlobal = min(minGlobal, minCur), max(maxGlobal, maxCur)
	}
	if total == minGlobal {
		return maxGlobal
	}
	return max(maxGlobal, total-minGlobal)
}

// func main() {
// 	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
// 	fmt.Println(maxSubarraySumCircular([]int{-5, -2, -5}))
// 	fmt.Println(maxSubarraySumCircular([]int{5, -2, 5}))
// 	fmt.Println(maxSubarraySumCircular([]int{1, -1, 3, -2, 5}))
// }
