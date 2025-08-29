package main

func lengthOfLongestSubstring(s string) int {
	longest, nonRepetative := 0, map[byte]int{}

	if len(s) <= 1 {
		return len(s)
	}

	j := 0
	for ; j < len(s); j++ {
		index, exists := nonRepetative[s[j]]
		if exists {
			longest = max(len(nonRepetative), longest)
			delete(nonRepetative, s[j])
			j = index
			continue
		}

		nonRepetative[s[j]] = j
	}
	longest = max(len(nonRepetative), longest)

	return longest
}

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
// 	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
// 	fmt.Println(lengthOfLongestSubstring("pwewkew"))  // 3
// 	fmt.Println(lengthOfLongestSubstring("zxyzxyz"))  // 3
// 	fmt.Println(lengthOfLongestSubstring("xxxx"))     // 1
// 	fmt.Println(lengthOfLongestSubstring("x"))        // 1
// 	fmt.Println(lengthOfLongestSubstring(""))         // 0
// 	fmt.Println(lengthOfLongestSubstring("au"))       // 2
// 	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
// 	fmt.Println(lengthOfLongestSubstring("dvczdf"))   // 5
// }
