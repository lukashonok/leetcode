package main

import "fmt"

type MinMaxQueue struct {
	Q []int
	D []int
}

func (queue *MinMaxQueue) enqueElement(element int) {
	// If there is no element in the queue
	if len(queue.Q) == 0 {
		queue.Q = append(queue.Q, element)
		queue.D = append(queue.D, element)
	} else {
		queue.Q = append(queue.Q, element)

		// Pop the elements out until the element at back is greater than current element
		for len(queue.D) > 0 && queue.D[len(queue.D)-1] < element {
			queue.D = queue.D[:len(queue.D)-1]
		}

		queue.D = append(queue.D, element)
	}
}

// Function to pop the element out from the queue
func (queue *MinMaxQueue) dequeElement() {
	// Condition when the Minimum element is the element at the front of the Deque
	if queue.Q[0] == queue.D[0] {
		queue.Q = queue.Q[1:]
		queue.D = queue.D[1:]
	} else {
		queue.Q = queue.Q[1:]
	}
}

// Function to pop the element out from the queue
func (queue *MinMaxQueue) getMax() int {
	return queue.D[0]
}

func maxSlidingWindowOld(nums []int, k int) []int {
	maxWindow := make([]int, len(nums)-k+1)

	queue := MinMaxQueue{}

	for i := range k {
		queue.enqueElement(nums[i])
	}

	i := 0
	fmt.Println("------------")
	fmt.Println("Q - ", queue.Q)
	fmt.Println("D - ", queue.D)
	fmt.Println("------------")
	for i < len(nums)-k {
		maxWindow[i] = queue.getMax()
		queue.dequeElement()
		fmt.Println("------------")
		fmt.Println("Q - ", queue.Q)
		fmt.Println("D - ", queue.D)
		fmt.Println("------------")
		queue.enqueElement(nums[i+k])
		fmt.Println("------------")
		fmt.Println("Q - ", queue.Q)
		fmt.Println("D - ", queue.D)
		fmt.Println("------------")
		i++
	}

	maxWindow[i] = queue.getMax()

	return maxWindow
}

// func main() {
// 	// fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) // [3,3,5,5,6,7]
// 	// fmt.Println(maxSlidingWindow([]int{1}, 1))                        // [1]
// 	// fmt.Println(maxSlidingWindow([]int{1, 2, 1, 0, 4, 2, 6}, 3))      // [2,2,4,4,6]
// 	fmt.Println(maxSlidingWindowOld([]int{1, 3, 1, 2, 0, 5}, 3)) // [3,3,2,5]
// }
