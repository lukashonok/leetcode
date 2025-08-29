package main

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 0; i < l1; i++ {
		dp[i][l2] = l1 - i
	}
	for j := 0; j < l2; j++ {
		dp[l1][j] = l2 - j
	}

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = 1 + min(dp[i+1][j], dp[i][j+1], dp[i+1][j+1])
			}
		}
	}

	return dp[0][0]
}

// func main() {
// 	fmt.Println(minDistance("abc", "acd"))
// }
