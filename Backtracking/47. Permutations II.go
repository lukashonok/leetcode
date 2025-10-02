package main

func permuteUnique(nums []int) [][]int {
	result, solution, length, used := [][]int{}, []int{}, len(nums), map[int]int{}

	for _, n := range nums {
		used[n]++
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i == length {
			currentSolution := make([]int, len(solution))
			copy(currentSolution, solution)
			result = append(result, currentSolution)
			return
		}

		for num, count := range used {
			if count == 0 {
				continue
			}

			used[num]--
			solution = append(solution, num)
			dfs(i + 1)
			solution = solution[:len(solution)-1]
			used[num]++
		}
	}
	dfs(0)

	return result
}

// func main() {
// 	fmt.Println(permuteUnique([]int{1, 1, 2}))
// }
