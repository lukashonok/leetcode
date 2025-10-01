package main

import "fmt"

type MinStackElement struct {
	Value int
	Next  *MinStackElement
}

type MinStack struct {
	Head *MinStackElement
	Min  *MinStack
}

func Constructor2() MinStack {
	return MinStack{
		Min: &MinStack{},
	}
}

func (this *MinStack) Push(val int) {
	if this.Min != nil {
		newMinValue := val
		if this.Head != nil && this.Min.Top() < val {
			newMinValue = this.Min.Top()
		}
		this.Min.Push(newMinValue)
	}

	oldHead := this.Head
	this.Head = &MinStackElement{
		Value: val,
		Next:  oldHead,
	}
}

func (this *MinStack) Pop() {
	if this.Min != nil {
		this.Min.Pop()
	}
	newHead := this.Head.Next
	this.Head = newHead
}

func (this *MinStack) Top() int {
	return this.Head.Value
}

func (this *MinStack) GetMin() int {
	return this.Min.Top()
}

func printMinStack(minStackElement *MinStackElement, tabs int) {
	totalTabs := "|"
	for i := 0; i < tabs; i++ {
		totalTabs += string("_")
	}
	fmt.Println(totalTabs, minStackElement.Value)
	if minStackElement.Next != nil {
		printMinStack(minStackElement.Next, tabs+1)
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// func main() {
// 	obj := Constructor()
// 	obj.Push(-2)
// 	obj.Push(0)
// 	obj.Push(-1)

// 	fmt.Println("Full")
// 	printStack(obj.Head, 0)
// 	fmt.Println("-------------")
// 	fmt.Println("Min")
// 	printStack(obj.Min.Head, 0)

// 	obj.Pop()
// 	// param_3 := obj.Top()
// 	// param_4 := obj.GetMin()

// }
