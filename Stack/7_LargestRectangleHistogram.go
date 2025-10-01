package main

type LRHStackElement struct {
	Value int
	Next  *LRHStackElement
}

type LRHStack struct {
	Head *LRHStackElement
}

func (stack *LRHStack) Push(val int) {
	oldHead := stack.Head
	stack.Head = &LRHStackElement{Value: val, Next: oldHead}
}

func (stack *LRHStack) Pop() {
	newHead := stack.Head.Next
	stack.Head = newHead
}

func (stack *LRHStack) Peek() int {
	return stack.Head.Value
}

func largestRectangleArea(heights []int) int {
	largestArea := 0
	stack := LRHStack{}

	lenHeight := len(heights)
	rightHeights := make([]int, lenHeight)
	for i := len(heights) - 1; i >= 0; i-- {
		for stack.Head != nil && heights[stack.Peek()] >= heights[i] {
			stack.Pop()
		}

		if stack.Head == nil {
			rightHeights[i] = -1
		} else {
			rightHeights[i] = stack.Peek() - 1
		}

		stack.Push(i)
	}

	stack = LRHStack{}
	leftHeights := make([]int, lenHeight)
	for i := 0; i < lenHeight; i++ {
		for stack.Head != nil && heights[stack.Peek()] >= heights[i] {
			stack.Pop()
		}

		if stack.Head == nil {
			leftHeights[i] = -1
		} else {
			leftHeights[i] = stack.Peek() + 1
		}

		stack.Push(i)
	}

	for i := range heights {
		if heights[i] == 0 {
			continue
		}
		if rightHeights[i] == -1 {
			rightHeights[i] = lenHeight - 1
		}
		if leftHeights[i] == -1 {
			leftHeights[i] = 0
		}
		currentArea := heights[i] * (rightHeights[i] - leftHeights[i] + 1)
		if currentArea > largestArea {
			largestArea = currentArea
		}
	}

	return largestArea
}

// func main() {
// 	start := time.Now()
// 	fmt.Println(largestRectangleArea([]int{1, 0, 5, 6, 2, 0}))
// 	fmt.Println(largestRectangleArea([]int{0, 9})) // 9
// 	fmt.Println(largestRectangleArea([]int{2, 4})) // 4
// 	fmt.Println("Total ms: ", time.Since(start).Milliseconds())
// }
