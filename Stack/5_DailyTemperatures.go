package main

type DTStackElement struct {
	Value int
	Next  *DTStackElement
}

type DTStack struct {
	Head *DTStackElement
}

func (this *DTStack) Push(val int) {
	oldHead := this.Head
	this.Head = &DTStackElement{Value: val, Next: oldHead}
}

func (this *DTStack) Pop() {
	newHead := this.Head.Next
	this.Head = newHead
}

func (this *DTStack) Peek() int {
	return this.Head.Value
}

func dailyTemperatures(temperatures []int) []int {
	stack := DTStack{Head: &DTStackElement{}}
	result := make([]int, len(temperatures))

	for i := 1; i < len(temperatures); i++ {
		lastIndex := stack.Peek()
		for stack.Head != nil && temperatures[i] > temperatures[stack.Peek()] {
			result[stack.Peek()] = lastIndex - stack.Peek() + 1
			stack.Pop()
		}
		stack.Push(i)
	}

	return result
}

// func main() {
// 	start := time.Now()
// 	// fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
// 	// fmt.Println(dailyTemperatures([]int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}))
// 	elapsed := time.Since(start)
// 	fmt.Printf("Took %s\n", elapsed)
// 	// fmt.Println(dailyTemperatures([]int{30, 60, 90}))
// 	// fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
// }
