package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	res := 0
	for i, j := 0, len(people)-1; i <= j; res, j = res+1, j-1 {
		if people[j]+people[i] <= limit {
			i++
		}
	}
	return res
}

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}
