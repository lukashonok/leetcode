package main

import (
	"fmt"
)

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k
	for l < r {
		m := (l + r) / 2
		if x-arr[m] > arr[m+k]-x {
			l = m + 1
		} else {
			r = m
		}
	}
	return arr[l : l+k]
}

func main() {

	fmt.Println(findClosestElements([]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8}, 2, 2)) // 1, 3
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))                // 1, 2, 3, 4
	fmt.Println(findClosestElements([]int{1, 2, 3, 6, 7}, 3, 5))                // 3, 6, 7
	fmt.Println(findClosestElements([]int{1, 1, 2, 3, 4, 5}, 4, -1))            // 1, 1, 2, 3
	fmt.Println(findClosestElements([]int{1}, 1, 1))                            // 1
	fmt.Println(findClosestElements([]int{1, 1, 1, 10, 10, 10}, 1, 9))          // 10
}
