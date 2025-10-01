package main

import (
	"strconv"
)

func calPoints(operations []string) int {
	s := []int{}
	for _, op := range operations {
		num, err := strconv.Atoi(op)
		if err == nil {
			s = append(s, num)
			continue
		}
		l := len(s)
		switch op {
		case "+":
			s = append(s, s[l-1]+s[l-2])
		case "C":
			s = s[:l-1]
		case "D":
			s = append(s, s[l-1]*2)
		}
	}
	res := 0
	for _, v := range s {
		res += v
	}
	return res
}

// func main() {
// 	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
// 	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
// 	fmt.Println(calPoints([]string{"1", "C"}))
// }
