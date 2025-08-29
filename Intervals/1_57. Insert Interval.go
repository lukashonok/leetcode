package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	} else if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	} else if newInterval[0] > intervals[len(intervals)-1][1] {
		return append(intervals, newInterval)
	}

	fixedIntervals := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if newInterval != nil && newInterval[0] < intervals[i][0] {
			fixedIntervals = append(fixedIntervals, newInterval)
			newInterval = nil
		}
		fixedIntervals = append(fixedIntervals, intervals[i])
	}
	if newInterval != nil {
		fixedIntervals = append(fixedIntervals, newInterval)
	}

	result, prev := [][]int{}, fixedIntervals[0]
	for i := 1; i < len(fixedIntervals); i++ {
		cur := fixedIntervals[i]
		if prev[1] >= cur[0] {
			prev[1] = max(cur[1], prev[1])
		} else {
			result = append(result, prev)
			prev = cur
		}
	}
	result = append(result, prev)

	return result
}

// func main() {
// 	// fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
// 	// fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
// 	// fmt.Println(insert([][]int{{1, 3}, {4, 6}}, []int{2, 5}))
// 	// fmt.Println(insert([][]int{{1, 2}, {3, 5}, {9, 10}}, []int{6, 7}))
// 	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7}))
// }
