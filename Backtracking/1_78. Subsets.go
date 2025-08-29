package main

func subsets(nums []int) [][]int {
	result, solution, length := [][]int{}, []int{}, len(nums)

	var iteration func(i int)
	iteration = func(i int) {
		if i == length {
			currentSolution := make([]int, len(solution))
			copy(currentSolution, solution)
			result = append(result, currentSolution)
			return
		}

		iteration(i + 1)

		solution = append(solution, nums[i])

		iteration(i + 1)

		solution = solution[:len(solution)-1]
	}
	iteration(0)
	return result
}

// func main() {
// 	fmt.Println(subsets([]int{1, 2, 3}))
// 	fmt.Println(subsets([]int{0}))
// }
