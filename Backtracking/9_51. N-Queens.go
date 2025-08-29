package main

import "fmt"

func solveNQueens(n int) [][]string {
	result := [][]string{}
	usedCols, usedPosDiags, usedNegDiags := map[int]bool{}, map[int]bool{}, map[int]bool{}
	board := [][]byte{}
	for i := 0; i < n; i++ {
		board = append(board, []byte{})
		for j := 0; j < n; j++ {
			board[i] = append(board[i], '.')
		}
	}

	var dfs func(r int)
	dfs = func(r int) {
		if r == n {
			newBoard := []string{}
			for i := 0; i < n; i++ {
				newBoard = append(newBoard, "")
				for j := 0; j < n; j++ {
					newBoard[i] += string(board[i][j])
				}
			}
			result = append(result, newBoard)
		}

		for c := 0; c < n; c++ {
			if usedCols[c] || usedNegDiags[r-c] || usedPosDiags[r+c] {
				continue
			}
			usedCols[c] = true
			usedNegDiags[r-c] = true
			usedPosDiags[r+c] = true
			board[r][c] = 'Q'
			dfs(r + 1)
			board[r][c] = '.'
			delete(usedCols, c)
			delete(usedNegDiags, r-c)
			delete(usedPosDiags, r+c)
		}
	}
	dfs(0)

	return result
}

func main() {
	fmt.Println(solveNQueens(14))
}
