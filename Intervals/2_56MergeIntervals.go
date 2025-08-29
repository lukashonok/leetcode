package main

func qsort(a [][]int) [][]int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivotIndex := len(a) / 2
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		if a[i][0] < a[right][0] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}

func merge(intervals [][]int) [][]int {
	intervals = qsort(intervals)
	if len(intervals) == 1 {
		return intervals
	}

	merged, prev := [][]int{}, intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] >= cur[0] {
			// must merge
			prev = []int{prev[0], max(cur[1], prev[1])}
		} else {
			merged = append(merged, prev)
			prev = cur
		}
	}
	merged = append(merged, prev)
	return merged
}

// func main() {
// 	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1,6],[8,10],[15,18]]
// 	fmt.Println(merge([][]int{{2, 6}, {1, 3}, {8, 15}, {9, 12}}))  // [[1 6] [8 15]]
// 	fmt.Println(merge([][]int{{2, 6}, {1, 3}, {8, 15}, {9, 12}}))  // [[1,6],[8,10],[15,18]]
// }
