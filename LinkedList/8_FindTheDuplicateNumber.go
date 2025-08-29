package main

func findDuplicate(nums []int) int {
	tort, hare := nums[0], nums[0]

	for {
		tort = nums[tort]
		hare = nums[nums[hare]]
		if tort == hare {
			break
		}
	}

	targetTort := nums[0]
	for targetTort != tort {
		tort = nums[tort]
		targetTort = nums[targetTort]
	}

	return targetTort
}

// func main() {
// 	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
// 	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
// 	fmt.Println(findDuplicate([]int{3, 3, 3, 3, 3}))
// 	fmt.Println(findDuplicate([]int{1, 2, 1}))
// 	fmt.Println(findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}))
// }
