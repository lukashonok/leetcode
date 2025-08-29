package main

func combinationSum(candidates []int, target int) [][]int {
	result, solution, length := [][]int{}, []int{}, len(candidates)

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

		iteration(i+1, currentSum)

		solution = append(solution, candidates[i])

		iteration(i, currentSum+candidates[i])

		solution = solution[:len(solution)-1]
	}
	iteration(0, 0)
	return result
}

// func main() {
// 	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
// }
