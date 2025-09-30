package main

import "fmt"

func isPalindrome(s string) (bool, int, int) {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false, l, r
		}
		l++
		r--
	}
	return true, 0, len(s) - 1
}

func validPalindrome(s string) bool {
	ok, l, r := isPalindrome(s)
	if ok {
		return true
	}

	fmt.Println(s[l:r])
	fmt.Println(s[l+1 : r+1])

	ok, _, _ = isPalindrome(s[l:r])
	if ok {
		return true
	}

	ok, _, _ = isPalindrome(s[l+1 : r+1])
	if ok {
		return true
	}

	fmt.Println(l, r)

	return false
}

func main() {
	check(validPalindrome("aba"), true, "aba")
	check(validPalindrome("abca"), true, "abca")
	check(validPalindrome("abc"), false, "abc")
}
