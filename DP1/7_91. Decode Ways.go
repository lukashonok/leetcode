package main

import (
	"slices"
)

func lessThanSix(c byte) bool {
	return slices.Index([]byte{'0', '1', '2', '3', '4', '5', '6'}, c) != -1
}

func numDecodings(s string) int {
	m := map[int]int{len(s): 1}
	var dfs func(i int) int
	dfs = func(i int) int {
		if m[i] > 0 {
			return m[i]
		}
		if s[i] == '0' {
			return 0
		}
		res := dfs(i + 1)
		if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && lessThanSix(s[i+1]))) {
			res += dfs(i + 2)
		}
		m[i] = res
		return res
	}
	return dfs(0)
}

// func main() {
// 	fmt.Println(numDecodings("12"))
// 	fmt.Println(numDecodings("226"))
// 	fmt.Println(numDecodings("06"))
// }
