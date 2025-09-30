package main

func subarraySum(nums []int, k int) int {
	prefixMap, res, curSum := map[int]int{0: 1}, 0, 0
	for _, num := range nums {
		curSum += num
		if count, ok := prefixMap[curSum-k]; ok {
			res += count
		}
		prefixMap[curSum]++
	}
	return res
}

// func main() {
// 	fmt.Println(subarraySum([]int{1, -1, 0}, 0))      // 3
// 	fmt.Println(subarraySum([]int{1, 2, 1, 2, 1}, 3)) // 4
// 	fmt.Println(subarraySum([]int{1, 1, 1}, 2))       // 2
// 	fmt.Println(subarraySum([]int{1, 2, 3}, 3))       // 2
// 	fmt.Println(subarraySum([]int{1}, 1))             // 1
// 	fmt.Println(subarraySum([]int{1}, 0))             // 0
// 	fmt.Println(subarraySum([]int{2, 1, 1}, 2))       // 2

// }
