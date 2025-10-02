package main

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	dp := map[int]int{len(s): 0}

	trie := Constructor()
	for _, word := range dictionary {
		trie.AddWord(word)
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if _, ok := dp[i]; ok {
			return dp[i]
		}

		res := 1 + dfs(i+1)
		for j := i; j < len(s); j++ {
			if trie.Search(s[i : j+1]) {
				res = min(res, dfs(j+1))
			}
		}

		dp[i] = res
		return res
	}

	return dfs(0)
}

func main() {
	fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	fmt.Println(minExtraChar("sayhelloworld", []string{"hello", "world"}))
}
