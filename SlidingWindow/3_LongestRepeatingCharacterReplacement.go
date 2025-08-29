package main

func characterReplacement(s string, k int) int {
	maxLength := 0
	letters := map[byte]bool{}
	for i := range s {
		_, exists := letters[s[i]]
		if !exists {
			letters[s[i]] = true
		}
	}

	for l := range letters {
		localMax, i, j := 0, 0, 0
		left := k

		for j < len(s) {
			if s[j] == l {
				localMax++
			} else if left == 0 {
				if s[i] != l && i != j {
					left++
				}

				i++
				j = max(i, j)
				localMax = max(localMax-1, 0)
				continue
			} else {
				left--
				localMax++
			}

			maxLength = max(localMax, maxLength)
			j++
		}
	}

	return maxLength
}

// func main() {
// 	fmt.Println(characterReplacement("AABABBA", 1) == 4)     // 4
// 	fmt.Println(characterReplacement("ABAB", 4) == 4)        // 4
// 	fmt.Println(characterReplacement("XYYX", 2) == 4)        // 4
// 	fmt.Println(characterReplacement("AAABABB", 1) == 5)     // 5
// 	fmt.Println(characterReplacement("ABCCBA", 2) == 4)      // 4
// 	fmt.Println(characterReplacement("ABABCDABACA", 2) == 5) // 5
// 	fmt.Println(characterReplacement("A", 1) == 1)           // 1
// 	fmt.Println(characterReplacement("ABC", 0) == 1)         // 1
// 	fmt.Println(characterReplacement("ABABCDABACA", 0) == 1) // 1
// 	fmt.Println(characterReplacement("ABAA", 0) == 2)        // 2
// 	fmt.Println(characterReplacement("BAAAB", 2) == 5)       // 5
// }
