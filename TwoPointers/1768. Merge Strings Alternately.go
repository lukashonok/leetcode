package main

func mergeAlternately(word1 string, word2 string) string {
	res, i, minLength := "", 0, min(len(word1), len(word2))
	for ; i < minLength; i++ {
		res += string(word1[i])
		res += string(word2[i])
	}
	for ; i < len(word1); i++ {
		res += string(word1[i])
	}
	for ; i < len(word2); i++ {
		res += string(word2[i])
	}
	return res
}

// func main() {
// 	check(mergeAlternately("abc", "pqr"), "apbqcr", "abc pqr")
// 	check(mergeAlternately("ab", "pqrs"), "apbqrs", "ab pqrs")
// 	check(mergeAlternately("abcd", "pq"), "apbqcd", "abcd pq")
// 	fmt.Println("Good")
// }
