package main

type visitMap map[int]map[int]bool

func pacificAtlantic(heights [][]int) [][]int {
	result, directions, m, n := [][]int{}, [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}, len(heights), len(heights[0])
	pac, atl := visitMap{}, visitMap{}
	for i := 0; i < len(heights); i++ {
		pac[i], atl[i] = map[int]bool{}, map[int]bool{}
		for j := 0; j < len((heights)[0]); j++ {
			pac[i][j], atl[i][j] = false, false
		}
	}

	var dfs func(r, c int, visited visitMap, prevHeight int, atl bool)
	dfs = func(r, c int, visited visitMap, prevHeight int, atl bool) {
		if r < 0 || r == m || c < 0 || c == n || visited[r][c] || heights[r][c] < prevHeight {
			return
		}
		visited[r][c] = true
		for _, d := range directions {
			newI, newJ := r+d[0], c+d[1]
			dfs(newI, newJ, visited, heights[r][c], atl)
		}
	}

	for r := 0; r < m; r++ {
		dfs(r, 0, pac, heights[r][0], false)
		dfs(r, n-1, atl, heights[r][n-1], true)
	}

	for c := 0; c < n; c++ {
		dfs(0, c, pac, heights[0][c], false)
		dfs(m-1, c, atl, heights[m-1][c], true)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pac[i][j] && atl[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

// func main() {
// 	// fmt.Println(pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
// 	// fmt.Println(pacificAtlantic([][]int{{1}}))
// 	fmt.Println(pacificAtlantic([][]int{{2, 1}, {1, 2}}))

// }
