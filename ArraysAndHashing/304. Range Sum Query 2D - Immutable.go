package main

import "fmt"

type NumMatrix struct {
}

func Constructor(matrix [][]int) NumMatrix {
	prefixMatrix := [][]int{}
	for range matrix {
		prefixMatrix = append(prefixMatrix, make([]int, len(matrix)))
	}
	for i := range matrix {
		for j := range matrix[0] {
			left := 0
			if j-1 >= 0 {
				left = matrix[i][j-1]
			}
			prefixMatrix[i][j] = left
		}
	}
	for i := range matrix {
		for j := range matrix[0] {
			top := 0
			if i-1 >= 0 {
				top = prefixMatrix[i-1][j]
			}
			prefixMatrix[i][j] = top
		}
	}
	return NumMatrix{}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return 0
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 *
 *
 */

func main() {
	obj := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(obj)
	// param_1 := obj.SumRegion(row1, col1, row2, col2)
}
