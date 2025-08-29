package main

type pack map[byte]bool

func isValidSudoku(board [][]byte) bool {
	size := 9
	columnsPacks := make([]pack, size)
	rowsPacks := make([]pack, size)
	squarePacks := make([]pack, size)

	for i := 0; i < size; i++ {
		columnsPacks[i] = pack{}
		rowsPacks[i] = pack{}
		squarePacks[i] = pack{}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// Row value check
			if board[i][j] != '.' {
				_, rowNumberExists := rowsPacks[i][board[i][j]]
				if rowNumberExists {
					return false
				}
				rowsPacks[i][board[i][j]] = true

				// Square value check
				squareIndex := (i/3)*3 + j/3
				_, squareNumberExists := squarePacks[squareIndex][board[i][j]]
				if squareNumberExists {
					return false
				}
				squarePacks[squareIndex][board[i][j]] = true
			}

			// Column value check
			if board[j][i] != '.' {
				_, columnNumberExists := columnsPacks[i][board[j][i]]
				if columnNumberExists {
					return false
				}
				columnsPacks[i][board[j][i]] = true
			}

		}
	}
	return true
}
