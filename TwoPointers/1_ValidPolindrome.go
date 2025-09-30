package main

import (
	"fmt"
	"strings"
)

func characterCheck(c byte) bool {
	return (c >= 48 && c <= 57) || (c >= 97 && c <= 122)
}

func check(got, want any, desc string) {
	if got != want {
		panic(fmt.Sprintf("FAIL: %s → got %v, want %v", desc, got, want))
	}
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	length := len(s)
	i := 0
	j := length - 1

	for i < j {
		if !characterCheck(s[i]) {
			i++
			continue
		}
		if !characterCheck(s[j]) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

// func main() {
// 	start := time.Now()
// 	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
// 	fmt.Println(isPalindrome("race a car"))                     // false
// 	fmt.Println(isPalindrome(" "))                              // true
// 	fmt.Println(isPalindrome("0P"))                             // false
// 	fmt.Println("Total ms: ", time.Since(start).Milliseconds())
// }
