package main

func exist(board [][]byte, word string) bool {
	result, used := false, map[int]map[int]bool{}
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < len(board); i++ {
		used[i] = map[int]bool{}
		for j := 0; j < len(board[0]); j++ {
			used[i][j] = false
		}
	}

	var dfs func(i, j, c int)
	dfs = func(i, j, c int) {
		if used[i][j] || board[i][j] != word[c] || result {
			return
		}

		if c == len(word)-1 {
			result = true
			return
		}

		used[i][j] = true
		for _, d := range directions {
			newI, newJ := i+d[0], j+d[1]
			if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) {
				dfs(newI, newJ, c+1)
			}
		}
		used[i][j] = false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, 0)
			if result {
				return true
			}
		}
	}

	return false
}

// func main() {
// 	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
// 	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
// 	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
// }
