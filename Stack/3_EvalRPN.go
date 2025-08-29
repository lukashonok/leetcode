package main

import (
	"strconv"
)

type RPNStackElement struct {
	Value string
	Next  *RPNStackElement
}

type RPNStack struct {
	Head *RPNStackElement
}

func (this *RPNStack) Push(val string) {
	oldHead := this.Head
	this.Head = &RPNStackElement{
		Value: val,
		Next:  oldHead,
	}
}

func (this *RPNStack) Pop() {
	newHead := this.Head.Next
	this.Head = newHead
}

func (this *RPNStack) Peek() string {
	return this.Head.Value
}

func evalRPNIteration(stack *RPNStack) int {
	operation := stack.Peek()
	stack.Pop()

	operation_number, err := strconv.Atoi(operation)
	if err == nil {
		return operation_number
	}

	part_second := evalRPNIteration(stack)

	part_first := evalRPNIteration(stack)

	result := 0
	switch operation {
	case "+":
		result = part_first + part_second
	case "-":
		result = part_first - part_second
	case "*":
		result = part_first * part_second
	case "/":
		result = part_first / part_second
	}

	return result
}

func evalRPN(tokens []string) int {
	stack := RPNStack{}

	for _, v := range tokens {
		stack.Push(v)
	}

	return evalRPNIteration(&stack)
}

// func printRPNStack(RPNStackElement *RPNStackElement, tabs int) {
// 	totalTabs := "|"
// 	for i := 0; i < tabs; i++ {
// 		totalTabs += string("_")
// 	}
// 	fmt.Println(totalTabs, RPNStackElement.Value)
// 	if RPNStackElement.Next != nil {
// 		printRPNStack(RPNStackElement.Next, tabs+1)
// 	}
// }

// func main() {
// 	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
// 	// fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
// 	// fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
// }
