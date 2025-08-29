package main

import "fmt"

type group [2]int

func maxCoins(nums []int) int {
	cache, n := map[group]int{}, len(nums)

	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if _, ok := cache[group{l, r}]; ok {
			return cache[group{l, r}]
		}

		cache[group{l, r}] = 0

		maxLocal := 0
		for i := l; i < r+1; i++ {
			left, cur, right := 1, nums[i], 1
			if l-1 >= 0 {
				left = nums[l-1]
			}
			if r+1 <= n-1 {
				right = nums[r+1]
			}
			coins := left*cur*right + dfs(l, i-1) + dfs(i+1, r)
			if (coins) > maxLocal {
				maxLocal = coins
			}
		}

		cache[group{l, r}] = maxLocal

		return cache[group{l, r}]
	}

	return dfs(0, n-1)
}

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
