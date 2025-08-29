package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tort, hare := head, head

	for i := 0; i < n; i++ {
		hare = hare.Next
	}

	for {
		if hare == nil || hare.Next == nil {
			break
		}
		tort, hare = tort.Next, hare.Next
	}

	if head == tort && hare == nil {
		head = head.Next
	} else if tort.Next.Next != nil {
		tort.Next = tort.Next.Next
	} else {
		tort.Next = nil
	}

	return head
}

// func main() {
// 	removeNthFromEnd(populateList([]int{2, 4, 6, 8, 10}), 2).Print() // 2, 4, 6, 10
// 	removeNthFromEnd(populateList([]int{2, 4, 6, 8, 10}), 1).Print() // 2, 4, 6, 8
// 	removeNthFromEnd(populateList([]int{2, 4, 6, 8, 10}), 5).Print() // 4, 6, 8, 10
// 	removeNthFromEnd(populateList([]int{1, 2}), 2).Print()           // 2
// 	removeNthFromEnd(populateList([]int{1, 2}), 1).Print()           // 1
// 	removeNthFromEnd(populateList([]int{2}), 1).Print()              //
// }
