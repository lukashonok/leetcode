package main

import (
	"math"
)

func kClosest(points [][]int, k int) [][]int {
	distances := []HeapData{}
	for i, point := range points {
		distances = append(distances, HeapData{key: float32(math.Sqrt(math.Pow(float64(point[0]), 2) + math.Pow(float64(point[1]), 2))), value: i})
	}
	MinHeapBuild(distances)
	result := [][]int{}
	for i := 0; i < k; i++ {
		heapPoint := MinHeapPop(&distances)
		result = append(result, points[heapPoint.value.(int)])
	}
	return result
}

// func main() {
// 	fmt.Println(kClosest(
// 		[][]int{{3, 3}, {5, -1}, {-2, 4}},
// 		2,
// 	))
// 	fmt.Println(kClosest(
// 		[][]int{{1, 3}, {-2, 2}},
// 		1,
// 	))
// }
