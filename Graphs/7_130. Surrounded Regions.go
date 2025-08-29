package main

import "fmt"

func solve(board [][]byte) {
	untouchableMap, visited, m, n := visitMap{}, visitMap{}, len(board), len(board[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < m; i++ {
		untouchableMap[i], visited[i] = map[int]bool{}, map[int]bool{}
		for j := 0; j < n; j++ {
			untouchableMap[i][j], visited[i][j] = false, false
		}
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r == m || c < 0 || c == n || visited[r][c] {
			return
		}

		visited[r][c] = true

		if board[r][c] == 'X' {
			return
		}

		untouchableMap[r][c] = true

		for _, d := range directions {
			dfs(r+d[0], c+d[1])
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !untouchableMap[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

// func main() {
// 	a := [][]byte{
// 		{'X', 'X', 'X', 'X'},
// 		{'X', 'O', 'O', 'X'},
// 		{'X', 'X', 'O', 'X'},
// 		{'X', 'O', 'X', 'X'},
// 	}
// 	solve(a)
// 	printGrid(a)
// 	a = [][]byte{
// 		{'X', 'X', 'X', 'X'},
// 		{'O', 'O', 'O', 'X'},
// 		{'X', 'X', 'O', 'X'},
// 		{'X', 'O', 'X', 'X'},
// 	}
// 	solve(a)
// 	printGrid(a)
// 	a = [][]byte{
// 		{'X', 'O', 'X', 'O', 'X', 'O'},
// 		{'O', 'X', 'O', 'X', 'O', 'X'},
// 		{'X', 'O', 'X', 'O', 'X', 'O'},
// 		{'O', 'X', 'O', 'X', 'O', 'X'},
// 	}
// 	solve(a)
// 	printGrid(a)
// }
