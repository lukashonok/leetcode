package main

import "fmt"

func swimInWater(grid [][]int) int {
	n := len(grid)
	heap := []HeapData{{grid[0][0], 0, 0}}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := map[[2]int]bool{}
	for len(heap) != 0 {
		cur := MinHeapPop(&heap)

		if visited[[2]int{cur[1], cur[2]}] {
			continue
		}

		visited[[2]int{cur[1], cur[2]}] = true
		if cur[1] == n-1 && cur[2] == n-1 {
			return cur[0]
		}
		for _, dir := range dirs {
			newI, newJ := cur[1]+dir[0], cur[2]+dir[1]
			if newI >= 0 && newJ >= 0 && newI < n && newJ < n {
				MinHeapInsert(&heap, HeapData{max(grid[newI][newJ], cur[0]), newI, newJ})
			}
		}

	}
	return 0
}

func main() {
	fmt.Println(swimInWater([][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}))
	fmt.Println(swimInWater([][]int{{0, 2}, {1, 3}}))
}
