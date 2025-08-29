package main

type pos [2]int

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	cache := map[pos]int{}

	var dfs func(i, j, prev int) int
	dfs = func(i, j, prev int) int {
		if i < 0 || i == m || j < 0 || j == n || matrix[i][j] <= prev {
			return 0
		}
		if _, ok := cache[pos{i, j}]; ok {
			return cache[pos{i, j}]
		}

		res := 1
		for _, dir := range dirs {
			newI, newJ := i+dir[0], j+dir[1]
			res = max(res, 1+dfs(newI, newJ, matrix[i][j]))
		}

		cache[pos{i, j}] = res
		return res
	}

	longestPath := 0
	for i := range m {
		for j := range n {
			longestPath = max(longestPath, dfs(i, j, -1))
		}
	}

	return longestPath
}

// func main() {
// 	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
// }
