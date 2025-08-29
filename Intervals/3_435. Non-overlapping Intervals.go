package main

import (
	"slices"
)

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a []int, b []int) int {
		if a[0] == b[0] {
			return 0
		} else if a[0] > b[0] {
			return 1
		} else {
			return -1
		}
	})

	lastEnd, result := intervals[0][1], 0
	for _, interval := range intervals[1:] {
		start, end := interval[0], interval[1]
		if start >= lastEnd {
			lastEnd = end
		} else {
			result++
			lastEnd = min(lastEnd, end)
		}
	}

	return result
}

// func main() {
// 	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 4}, {1, 4}})) //1
// 	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 4}}))         // 0
// 	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 4}}))         // 0
// }
