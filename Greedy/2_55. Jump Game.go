package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	goal := len(nums) - 1
	for i := goal - 1; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	return goal == 0
}

// func main() {
// 	// fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
// 	fmt.Println(canJump([]int{1, 2, 1, 0, 1}))
// }
