package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	var minIndexQueue *ListNode

	enqueMinElement := func(index int) {
		if lists[index] == nil {
			return
		}

		cur := minIndexQueue
		var last *ListNode
		newElement := &ListNode{Val: index}
		if cur == nil {
			minIndexQueue = newElement
			return
		}

		for cur != nil {
			if lists[index].Val < lists[cur.Val].Val {
				// insert before
				newElement.Next = cur
				if last != nil {
					last.Next = newElement
				} else {
					minIndexQueue = newElement
				}
				return
			}
			last = cur
			cur = cur.Next
		}
		// insert after
		last.Next = newElement
	}

	for i := 0; i < len(lists); i++ {
		enqueMinElement(i)
	}

	var newHead *ListNode
	var curHead *ListNode
	for minIndexQueue != nil {
		minElementIndex := minIndexQueue.Val
		minValue := &ListNode{Val: lists[minElementIndex].Val}
		lists[minElementIndex] = lists[minElementIndex].Next
		minIndexQueue = minIndexQueue.Next
		enqueMinElement(minElementIndex)

		if newHead == nil {
			newHead = minValue
			curHead = newHead
		} else {
			curHead.Next = minValue
			curHead = curHead.Next
		}

	}

	return newHead
}
