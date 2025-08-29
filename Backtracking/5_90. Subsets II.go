package main

import (
	"slices"
)

func subsetsWithDup(nums []int) [][]int {
	result, solution, length := [][]int{}, []int{}, len(nums)
	slices.Sort(nums)

	var iteration func(i int)
	iteration = func(i int) {
		if i == length {
			currentSolution := make([]int, len(solution))
			copy(currentSolution, solution)
			result = append(result, currentSolution)
			return
		}

		// Here we deside to add everything
		solution = append(solution, nums[i])
		iteration(i + 1)
		solution = solution[:len(solution)-1]

		// Here we deside to not add same number again, skippin this to next number
		for i+1 < length && nums[i] == nums[i+1] {
			i++
		}
		iteration(i + 1)
	}
	iteration(0)
	return result
}

// func main() {
// 	fmt.Println(subsetsWithDup([]int{1, 2, 2, 3}))
// }
