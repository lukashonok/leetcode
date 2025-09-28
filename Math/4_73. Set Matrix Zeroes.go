package main

func setZeroes(matrix [][]int) {
	m, n, rowZero := len(matrix), len(matrix[0]), false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				if i > 0 {
					matrix[i][0] = 0
				} else {
					rowZero = true
				}
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if rowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}

// func main() {
// 	a := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
// 	setZeroes(a)
// 	fmt.Println(a)
// 	a = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
// 	setZeroes(a)
// 	fmt.Println(a)
// }
