package main

type ptrs [2]int

func numDistinct(s string, t string) int {
	cache, m, n := map[ptrs]int{}, len(s), len(t)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j == n {
			return 1
		}
		if i == m {
			return 0
		}
		if _, ok := cache[ptrs{i, j}]; ok {
			return cache[ptrs{i, j}]
		}
		if s[i] == t[j] {
			cache[ptrs{i, j}] = dfs(i+1, j+1) + dfs(i+1, j)
		} else {
			cache[ptrs{i, j}] = dfs(i+1, j)
		}

		return cache[ptrs{i, j}]
	}

	return dfs(0, 0)
}

// func main() {
// 	fmt.Println(numDistinct("babgbag", "bag"))
// }
