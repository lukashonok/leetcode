package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	stack := []string{}
	for _, c := range s {
		if c != ']' {
			stack = append(stack, string(c))
			continue
		}

		currentPart := ""
		for stack[len(stack)-1] != "[" {
			currentPart = string(stack[len(stack)-1]) + currentPart
			stack = stack[:len(stack)-1]
		}
		stack = stack[:len(stack)-1] // popping [
		numberRaw := ""
		for len(stack) > 0 {
			if _, err := strconv.Atoi(stack[len(stack)-1]); err == nil {
				numberRaw = string(stack[len(stack)-1]) + numberRaw
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		number, _ := strconv.Atoi(numberRaw)

		multiString := ""
		for i := 0; i < number; i++ {
			multiString += currentPart
		}

		stack = append(stack, multiString)
	}
	res := ""
	for len(stack) > 0 {
		res = stack[len(stack)-1] + res
		stack = stack[:len(stack)-1]
	}
	return res
}

func main() {

	// fmt.Println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")) // zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef
	fmt.Println(decodeString("xa3[a2[c]]")) // accaccacc
	// fmt.Println(decodeString("12[abcd]2[bc]")) // aaabcbc
	fmt.Println(decodeString("3[a]pp2[bc]"))   // aaabcbc
	fmt.Println(decodeString("2[abc]3[cd]ef")) // abcabccdcdcdef
}
