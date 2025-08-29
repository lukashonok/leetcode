package main

import "fmt"

func checkValidString(s string) bool {
	left, right := 0, 0
	for _, c := range s {
		if c == '(' {
			left++
			right++
		} else if c == ')' {
			left--
			right--
		} else if c == '*' {
			left--
			right++
		}
		if right < 0 {
			return false
		}
		if left < 0 {
			left = 0
		}
	}
	return left == 0
}

func main() {
	fmt.Println(checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())")) // false

	fmt.Println(checkValidString("()"))         // true
	fmt.Println(checkValidString("(*)"))        // true
	fmt.Println(checkValidString("(*))"))       // true
	fmt.Println(checkValidString(")"))          // true
	fmt.Println(checkValidString("())"))        // true
	fmt.Println(checkValidString("*())"))       // true
	fmt.Println(checkValidString("****())"))    // true
	fmt.Println(checkValidString("****())())")) // true

}
