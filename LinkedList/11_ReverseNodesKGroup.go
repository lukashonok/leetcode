package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	start, last := head, head

	var reverseIteration func(*ListNode)
	reverseIteration = func(current *ListNode) {
		next := current.Next
		current.Next = start
		start = current
		if next != nil {
			reverseIteration(next)
		}
	}

	i := 1
	for cur := head; cur != nil; i = i + 1 {
		if i%k != 0 {
			cur = cur.Next
			continue
		}

		first := start
		next := cur.Next
		cur.Next = nil
		reverseIteration(start)
		if i == k {
			head = start
		} else {
			last.Next = start
		}
		first.Next = next
		last = first
		cur = next
		start = next
	}

	return head
}
