package main

func permute(nums []int) [][]int {
	result, solution, length, used := [][]int{}, []int{}, len(nums), map[int]bool{}

	var dfs func(i int)
	dfs = func(i int) {
		if len(solution) == length {
			currentSolution := make([]int, len(solution))
			copy(currentSolution, solution)
			result = append(result, currentSolution)
			return
		}

		for _, num := range nums {
			_, exists := used[num]
			if exists {
				continue
			}

			used[num] = true

			solution = append(solution, num)
			dfs(i + 1)
			solution = solution[:len(solution)-1]

			delete(used, num)
		}
	}
	dfs(0)

	return result
}

// func main() {
// 	fmt.Println(permute([]int{1, 2, 3}))
// }
