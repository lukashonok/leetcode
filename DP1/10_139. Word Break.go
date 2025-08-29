package main

func wordBreak(s string, wordDict []string) bool {
	length := len(s)
	dp := make([]bool, length+1)
	dp[length] = true

	for i := length - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if i+len(w) <= len(s) && s[i:i+len(w)] == w {
				dp[i] = dp[i+len(w)]
			}
			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}

// func main() {
// 	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
// }
