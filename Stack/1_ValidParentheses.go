package main

type StackElement struct {
	Value byte
	Next  *StackElement
}

type Stack struct {
	Head *StackElement
}

func stackPut(stack *Stack, value byte) {
	head := stack.Head
	newElement := StackElement{
		Value: value,
		Next:  head,
	}
	stack.Head = &newElement
}

func stackPop(stack *Stack) *StackElement {
	head := stack.Head
	stack.Head = head.Next
	return head
}

func isValid(s string) bool {
	stack := Stack{}
	parenthesesMap := map[byte]byte{'{': '.', '(': '.', '[': '.', '}': '{', ')': '(', ']': '['}

	for _, c := range s {
		parenthesis := byte(c)
		openedFlag := parenthesesMap[parenthesis]

		// if opened
		if openedFlag == '.' {
			stackPut(&stack, parenthesis)
			continue
		}

		// if closed more than opened
		if stack.Head == nil {
			return false
		}

		// closed check
		lastParenthesis := stackPop(&stack).Value
		openedParenthesis := parenthesesMap[parenthesis]
		if openedParenthesis != lastParenthesis {
			return false
		}
	}

	return stack.Head == nil
}
