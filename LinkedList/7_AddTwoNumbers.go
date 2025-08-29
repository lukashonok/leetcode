package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newNumber *ListNode
	var lastDigit *ListNode

	extra := 0

	for {
		firstDigit, secondDigit, alarm := 0, 0, 0
		if l1 != nil {
			firstDigit = l1.Val
			l1 = l1.Next
		} else {
			alarm++
		}
		if l2 != nil {
			secondDigit = l2.Val
			l2 = l2.Next
		} else {
			alarm++
		}
		if alarm == 2 {
			break
		}
		result := firstDigit + secondDigit + extra
		extra = 0
		if result >= 10 {
			result -= 10
			extra = 1
		}
		newDigit := &ListNode{Val: result}
		if newNumber == nil {
			lastDigit = newDigit
			newNumber = lastDigit
		} else {
			lastDigit.Next = newDigit
			lastDigit = newDigit
		}
	}

	if extra == 1 {
		lastDigit.Next = &ListNode{Val: 1}
	}

	return newNumber
}

// func main() {
// 	addTwoNumbers(populateList([]int{2, 4, 9}), populateList([]int{5, 6, 4, 9})).Print()
// 	addTwoNumbers(populateList([]int{8}), populateList([]int{7})).Print()
// 	addTwoNumbers(populateList([]int{2, 4, 3}), populateList([]int{5, 6, 4})).Print()
// 	addTwoNumbers(populateList([]int{0}), populateList([]int{0})).Print()
// 	addTwoNumbers(populateList([]int{9, 9, 9, 9, 9, 9, 9}), populateList([]int{9, 9, 9, 9})).Print()
// }
