package main

import "fmt"

func wordBreak2(s string, wordDict []string) bool {
	l := len(s)
	dp := make([]bool, l+1)
	dp[l] = true

	for i := l - 1; i >= 0; i-- {
		for _, word := range wordDict {
			lw := len(word)
			if i+lw <= l && s[i:i+lw] == word && dp[i+lw] {
				dp[i] = dp[i+lw]
			}
		}
	}

	return dp[0]
}

func main() {
	fmt.Println(wordBreak2("abcd", []string{"a", "abc", "b", "cd"}))
	fmt.Println(wordBreak2("leetcode", []string{"leet", "code"}))
}
