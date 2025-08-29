package main

import (
	"strconv"
)

func getNumber(s string) (int, int) {
	number, numberLength := "", 0
	for _, v := range s {
		if v > 71 {
			break
		}
		number += string(v)
		numberLength++
	}
	value, _ := strconv.Atoi(number)
	return value, numberLength
}

func decodeStringIteration(s string) (string, string) {
	charsBefore := ""
	i := 0
	for len(s) != 0 && i < len(s) && s[i] > 71 && s[i] != ']' {
		charsBefore += string(s[i])
		i++
	}
	s = s[i:]

	if len(s) == 0 {
		return charsBefore, s
	} else if s[0] == ']' {
		return charsBefore, s[1:]
	}

	number, i := getNumber(s)
	s = s[i:]
	duplicated := ""
	stringToDuplicate, s := decodeStringIteration(s[1:])

	for i := 0; i < number; i++ {
		duplicated += stringToDuplicate
	}

	return charsBefore + duplicated, s
}

func decodeString(s string) string {
	result, localResult := "", ""
	for s != "" {
		localResult, s = decodeStringIteration(s)
		result += localResult
	}
	return result
}

// func main() {
// 	fmt.Println(decodeString("3[a2[c]]2[ab]"))                // accaccaccabab
// 	fmt.Println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")) // zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef
// 	// fmt.Println(decodeString("3[a]2[bc]"))     // aaabcbc
// 	// fmt.Println(decodeString("3[a2[c]]"))      // accaccacc
// 	fmt.Println(decodeString("2[abc]3[cd]ef")) // abcabccdcdcdef
// }
