package main

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

// func main() {
// 	fmt.Println(getConcatenation([]int{1, 2, 1}))
// }
