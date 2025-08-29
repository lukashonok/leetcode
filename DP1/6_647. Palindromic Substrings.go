package main

func countSubstrings(s string) int {
	res := 0

	expandFromMiddle := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			r++
			l--
			res++
		}
	}

	for i := range len(s) {
		expandFromMiddle(i, i)
		expandFromMiddle(i, i+1)
	}

	return res
}

// func main() {
// 	fmt.Println(countSubstrings("abc"))
// 	fmt.Println(countSubstrings("babad"))
// 	fmt.Println(countSubstrings("aaa"))
// }
