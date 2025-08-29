package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Print() {
	if head == nil {
		fmt.Println("queue:  empty")
		return
	}
	last := head
	total := ""
	for last.Next != nil {
		total += fmt.Sprint(last.Val) + " "
		last = last.Next
	}
	total += fmt.Sprint(last.Val) + " "
	fmt.Println("queue: ", total)
}

func populateList(nums []int) *ListNode {
	head := ListNode{}

	if len(nums) == 0 {
		return &head
	}

	head.Val = nums[0]

	for i := 1; i < len(nums); i++ {
		Enque(&head, nums[i])
	}

	return &head
}

func Enque(head *ListNode, value int) {
	last := head
	for last.Next != nil {
		last = last.Next
	}
	newNode := ListNode{Val: value}
	last.Next = &newNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var reversedHead *ListNode
	var reverseIteration func(*ListNode)

	reverseIteration = func(current *ListNode) {
		next := current.Next
		current.Next = reversedHead
		reversedHead = current
		if next != nil {
			reverseIteration(next)
		}
	}

	reverseIteration(head)

	return reversedHead
}

// func main() {
// 	reverseList(populateList([]int{1, 2, 3})).Print()       // 3, 2, 1
// 	reverseList(populateList([]int{2, 1})).Print()          // 1,2
// 	reverseList(populateList([]int{1, 2, 3, 4, 5})).Print() // 5,4,3,2,1
// 	fmt.Println(reverseList(nil))
// }
