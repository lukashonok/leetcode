package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	dirs, currentDir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}, 0
	l, r, b, t := 0, n-1, m-1, 0
	corner := [2]int{t, r}
	i, j := 0, 0
	path := []int{matrix[0][0]}
	for len(path) != m*n {
		if [2]int{i, j} == corner {
			switch currentDir {
			case 0:
				t++
				corner = [2]int{b, r}
			case 1:
				r--
				corner = [2]int{b, l}
			case 2:
				b--
				corner = [2]int{t, l}
			case 3:
				l++
				corner = [2]int{t, r}
			}
			currentDir++
			if currentDir == len(dirs) {
				currentDir = 0
			}
		}
		i, j = i+dirs[currentDir][0], j+dirs[currentDir][1]
		path = append(path, matrix[i][j])
	}
	return path
}

// func main() {
// 	fmt.Println(spiralOrder([][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}))
// }
