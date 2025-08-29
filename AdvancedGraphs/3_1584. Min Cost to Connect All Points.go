package main

import (
	"math"
)

func minCostConnectPoints(points [][]int) int {
	adj := map[int][][]int{}
	for i, point := range points {
		adj[i] = [][]int{}
		for j, innerPoint := range points {
			if i == j {
				continue
			}

			distance := int(math.Abs(float64(innerPoint[1]-point[1])) + math.Abs(float64(innerPoint[0]-point[0])))
			adj[i] = append(adj[i], []int{distance, j})
		}
	}

	heap, visited, result := []HeapData{{0, 0}}, map[int]bool{}, 0
	for len(heap) != 0 {
		current := MinHeapPop(&heap)
		if visited[current[1]] {
			continue
		}
		visited[current[1]] = true
		result = result + current[0]
		for _, v := range adj[current[1]] {
			if visited[v[1]] {
				continue
			}

			MinHeapInsert(&heap, []int{v[0], v[1]})
		}
	}

	return result
}

// func main() {
// 	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
// }
