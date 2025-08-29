package main

// L(0, 6) = L[0] == L[6] ? (L[1, 5] + 2) : (max(L[0, 5], L[1, 6]))

func longestPalindrome(s string) string {
	res, maxLength := "", 0

	expandFromMiddle := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			newLength := r - l + 1
			if newLength > maxLength {
				maxLength = newLength
				res = s[l : r+1]
			}
			r++
			l--
		}
	}

	for i := range len(s) {
		expandFromMiddle(i, i)
		expandFromMiddle(i, i+1)
	}

	return res
}

// func main() {
// 	fmt.Println(longestPalindrome("BABCBAB"))
// 	fmt.Println(longestPalindrome("babad"))
// 	fmt.Println(longestPalindrome("ac"))
// }
