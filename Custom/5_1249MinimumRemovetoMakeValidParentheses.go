package main

type Pair struct {
	index int
	value rune
}

func minRemoveToMakeValid(s string) string {
	stack, toDelete := []Pair{}, []int{}

	for i, c := range s {
		if c == '(' && (len(stack) == 0 || stack[0].value != ')') {
			stack = append(stack, Pair{index: i, value: c})
		} else if c == ')' && len(stack) == 0 {
			toDelete = append(toDelete, i)
		} else if c == ')' && stack[0].value == '(' {
			stack = stack[1:]
		}
	}

	for len(stack) != 0 || len(toDelete) != 0 {
		if len(stack) != 0 && len(toDelete) != 0 {
			stackLast, toDeleteLast := stack[len(stack)-1].index, toDelete[len(toDelete)-1]
			if stackLast > toDeleteLast {
				s = s[:stackLast] + s[stackLast+1:]
				stack = stack[:len(stack)-1]
			} else {
				s = s[:toDeleteLast] + s[toDeleteLast+1:]
				toDelete = toDelete[:len(toDelete)-1]
			}
		} else if len(stack) != 0 {
			stackLast := stack[len(stack)-1].index
			s = s[:stackLast] + s[stackLast+1:]
			stack = stack[:len(stack)-1]
		} else if len(toDelete) != 0 {
			toDeleteLast := toDelete[len(toDelete)-1]
			s = s[:toDeleteLast] + s[toDeleteLast+1:]
			toDelete = toDelete[:len(toDelete)-1]
		}
	}

	return s
}

// func main() {
// 	// fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)")) // lee(t(c)o)de
// 	// fmt.Println(minRemoveToMakeValid("a)b(c)d"))       // ab(c)d
// 	fmt.Println(minRemoveToMakeValid("))((")) // ""
// }
