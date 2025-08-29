package main

import (
	"strings"
)

func generateParenthesis(n int) []string {
	stack := []string{}
	res := []string{}

	var backtrack func(int, int)
	backtrack = func(opened int, closed int) {
		if opened+closed == n*2 {
			res = append(res, strings.Join(stack, ""))
		}

		if opened < n {
			stack = append(stack, "(")
			backtrack(opened+1, closed)
			stack = stack[:len(stack)-1]
		}

		if closed < opened {
			stack = append(stack, ")")
			backtrack(opened, closed+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return res
}

// func main() {
// 	fmt.Println(generateParenthesis(4))
// }
