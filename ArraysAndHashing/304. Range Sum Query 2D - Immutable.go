package main

type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{prefix: [][]int{}}
	}
	n, m := len(matrix), len(matrix[0])
	prefix := make([][]int, n+1)
	for i := range prefix {
		prefix[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{prefix: prefix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	p := this.prefix
	return p[row2+1][col2+1] - p[row1][col2+1] - p[row2+1][col1] + p[row1][col1]
}

// func main() {
// 	obj := Constructor([][]int{
// 		{-4, -2},
// 		{-4, -2},
// 	})
// 	// obj := Constructor([][]int{
// 	// 	{1, 2, 3, 4},
// 	// 	{5, 6, 7, 8},
// 	// 	{9, 10, 11, 12},
// 	// 	{13, 14, 15, 16},
// 	// })

// 	fmt.Println(obj.SumRegion(0, 0, 0, 0))
// 	fmt.Println(obj.SumRegion(0, 0, 0, 1))
// 	fmt.Println(obj.SumRegion(0, 1, 0, 1))

// 	obj = Constructor([][]int{
// 		{3, 0, 1, 4, 2},
// 		{5, 6, 3, 2, 1},
// 		{1, 2, 0, 1, 5},
// 		{4, 1, 0, 1, 7},
// 		{1, 0, 3, 0, 5},
// 	})

// 	fmt.Println(obj.SumRegion(1, 1, 2, 2))
// 	fmt.Println(obj.SumRegion(1, 2, 2, 4))
// 	// param_1 := obj.SumRegion(row1, col1, row2, col2)
// }
