package main

import (
	"slices"
)

func combinationSum2(candidates []int, target int) [][]int {
	result, solution, length := [][]int{}, []int{}, len(candidates)
	slices.Sort(candidates)

	var iteration func(i int, currentSum int)
	iteration = func(i int, currentSum int) {
		if currentSum == target {
			currentSolution := make([]int, len(solution))
			copy(currentSolution, solution)
			result = append(result, currentSolution)
			return
		} else if currentSum > target || i == length {
			return
		}

		// Here if 1, 1, 1, 1, 2 -> all 1s will be added
		solution = append(solution, candidates[i])
		iteration(i+1, currentSum+candidates[i])
		solution = solution[:len(solution)-1]

		// Here we try to not add all of them
		for i+1 < length && candidates[i] == candidates[i+1] {
			i++
		}
		iteration(i+1, currentSum)
	}
	iteration(0, 0)
	return result
}

// func main() {
// 	fmt.Println(combinationSum2([]int{2, 5, 4, 1, 7}, 7))
// }
