package main

func combine(n int, k int) [][]int {
	globalRes := [][]int{}

	res := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if k == len(res) {
			currentSolution := make([]int, len(res))
			copy(currentSolution, res)
			globalRes = append(globalRes, currentSolution)
			return
		}

		for j := i; j <= n; j++ {
			res = append(res, j)
			dfs(j + 1)
			res = res[:len(res)-1]
		}
	}
	dfs(1)

	return globalRes
}

// func main() {
// 	fmt.Println(combine(4, 2))
// }
