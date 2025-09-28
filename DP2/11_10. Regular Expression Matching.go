package main

import "fmt"

type b [2]int

func isMatch(s string, p string) bool {
	l1, l2 := len(s), len(p)
	cache := map[b]bool{}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if _, ok := cache[b{i, j}]; ok {
			return cache[b{i, j}]
		}

		if i >= l1 && j >= l2 {
			return true
		} else if j >= l2 {
			return false
		}

		match := i < l1 && (s[i] == p[j] || p[j] == '.')

		if j+1 < l2 && p[j+1] == '*' {
			cache[b{i, j}] = dfs(i, j+2) || (match && dfs(i+1, j))
		} else if match {
			cache[b{i, j}] = dfs(i+1, j+1)
		}

		return cache[b{i, j}]
	}
	return dfs(0, 0)
}

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
