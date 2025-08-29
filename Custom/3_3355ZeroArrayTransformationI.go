package main

import (
	"strconv"
	"strings"
)

func isZeroArray(nums []int, queries [][]int) bool {
	positionCounters := make([]int, len(nums))

	duplicates := map[string]int{}
	for i := 0; i < len(queries); i++ {
		duplicates[strconv.Itoa(queries[i][0])+"."+strconv.Itoa(queries[i][1])]++
	}

	for numberRange, value := range duplicates {
		parts := strings.Split(numberRange, ".")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])

		for i := from; i <= to; i++ {
			positionCounters[i] += value
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > positionCounters[i] {
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}}))
// 	fmt.Println(isZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3}, {0, 2}}))
// }
