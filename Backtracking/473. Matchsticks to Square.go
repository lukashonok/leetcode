package main

import (
	"sort"
)

func makesquare(matchsticks []int) bool {
	perimeter := 0
	for _, v := range matchsticks {
		perimeter += v
	}

	if perimeter%4 != 0 {
		return false
	}

	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	side := perimeter / 4

	sides := make([]int, 4)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i == len(matchsticks) {
			return true
		}

		for j := 0; j < 4; j++ {
			if sides[j]+matchsticks[i] <= side {
				sides[j] += matchsticks[i]
				if dfs(i + 1) {
					return true
				}
				sides[j] -= matchsticks[i]
			}
		}

		return false
	}

	return dfs(0)
}

// func main() {
// 	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
// 	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
// }
