package main

func rotate(matrix [][]int) {
	m, n := matrix, len(matrix)
	left, right := 0, n-1

	for left < right {
		for i := range right - left {
			top, bottom := left, right
			topLeft := m[top][left+i]
			m[top][left+i] = m[bottom-i][left]
			m[bottom-i][left] = m[bottom][right-i]
			m[bottom][right-i] = m[top+i][right]
			m[top+i][right] = topLeft
		}
		right--
		left++
	}
}

// func main() {
// 	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// 	rotate(m)
// 	fmt.Println(m)
// }
