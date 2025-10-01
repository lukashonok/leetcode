package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := head
	var nodeBefore, nodeAfter, reverseStart, reverseStartCopy, reverseEnd *ListNode
	cur := head

	i := 1
	for cur != nil {
		if cur.Next != nil && i+1 == left {
			nodeBefore = cur
			reverseStart = cur.Next
			reverseStartCopy = reverseStart
		}
		if reverseStart == nil && i == left {
			reverseStart = cur
			reverseStartCopy = reverseStart
		}
		if i == right {
			nodeAfter = cur.Next
			reverseEnd = cur
		}
		cur = cur.Next
		i++
	}

	if reverseEnd != nil {
		reverseEnd.Next = nil
	}

	var reverseIteration func(node *ListNode)
	reverseIteration = func(node *ListNode) {
		next := node.Next
		node.Next = reverseStart
		reverseStart = node

		if next != nil {
			reverseIteration(next)
		}
	}
	if reverseStart != nil {
		reverseIteration(reverseStart)
	}

	reverseStartCopy.Next = nodeAfter
	if nodeBefore == nil {
		newHead = reverseStart
	} else {
		nodeBefore.Next = reverseStart
	}

	return newHead
}

func main() {
	reverseBetween(populateList([]int{1, 2, 3, 4, 5}), 2, 4).Print()
	reverseBetween(populateList([]int{1, 2, 3, 4, 5}), 1, 5).Print()
	reverseBetween(populateList([]int{1, 2, 3, 4, 5}), 3, 4).Print()
	reverseBetween(populateList([]int{1, 2, 3, 4, 5}), 1, 3).Print()
	reverseBetween(populateList([]int{5}), 1, 1).Print()
	reverseBetween(populateList([]int{1, 2}), 1, 2).Print()
}
