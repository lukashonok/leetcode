package main

import "fmt"

type DequeElement struct {
	Value int
	Next  *DequeElement
	Prev  *DequeElement
}

type Deque struct {
	Head *DequeElement
	Tail *DequeElement
}

func (deque *Deque) Unshift(value int) {
	newElement := &DequeElement{Value: value}
	if deque.Head == nil {
		deque.Head = newElement
		deque.Tail = deque.Head
		return
	}

	deque.Tail.Next = newElement
	newElement.Prev = deque.Tail
	deque.Tail = newElement
}

func (deque *Deque) Shift() *int {
	if deque.Head == nil {
		return nil
	}
	oldHead := deque.Head
	deque.Head = deque.Head.Next
	if deque.Head == nil {
		deque.Tail = nil
	} else {
		deque.Head.Prev = nil
	}
	return &oldHead.Value
}

func (deque *Deque) Pop() *int {
	if deque.Tail == nil {
		return nil
	}
	oldTail := deque.Tail
	deque.Tail = deque.Tail.Prev
	if deque.Tail == nil {
		deque.Head = nil
	} else {
		deque.Tail.Next = nil
	}
	return &oldTail.Value
}

type MinMaxDeque struct {
	Q Deque
	D Deque
}

func (queue *MinMaxDeque) enqueElement(element int) {
	// If there is no element in the queue
	if queue.Q.Head == nil {
		queue.Q.Unshift(element)
		queue.D.Unshift(element)
	} else {
		queue.Q.Unshift(element)

		// Shift the elements out until the element at back is greater than current element
		for queue.D.Head != nil && queue.D.Tail.Value < element {
			queue.D.Pop()
		}

		queue.D.Unshift(element)
	}
}

// Function to Shift the element out from the queue
func (queue *MinMaxDeque) dequeElement() {
	// Condition when the Minimum element is the element at the front of the Deque
	if queue.Q.Head.Value == queue.D.Head.Value {
		queue.Q.Shift()
		queue.D.Shift()
	} else {
		queue.Q.Shift()
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	maxWindow := make([]int, len(nums)-k+1)

	queue := MinMaxDeque{}

	for i := range k {
		queue.enqueElement(nums[i])
	}

	i := 0
	for i < len(nums)-k {
		maxWindow[i] = queue.D.Head.Value
		queue.dequeElement()
		queue.enqueElement(nums[i+k])
		i++
	}

	maxWindow[i] = queue.D.Head.Value

	return maxWindow
}

func printDeque(queue MinMaxDeque) {
	fmt.Println("------------")
	total := ""
	for currentElement := queue.Q.Head; currentElement != nil; currentElement = currentElement.Next {
		total += fmt.Sprint(currentElement.Value) + " "
	}
	fmt.Println("Q - ", total)

	total = ""
	for currentElement := queue.D.Head; currentElement != nil; currentElement = currentElement.Next {
		total += fmt.Sprint(currentElement.Value) + " "
	}
	fmt.Println("D - ", total)

	fmt.Println("------------")
}

// func main() {
// 	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) // [3,3,5,5,6,7]
// 	fmt.Println(maxSlidingWindow([]int{1}, 1))                        // [1]
// 	fmt.Println(maxSlidingWindow([]int{1, 2, 1, 0, 4, 2, 6}, 3))      // [2,2,4,4,6]
// 	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))         // [3,3,2,5]
// }
