package main

// L0 → L1 → … → Ln - 1 → Ln
// L0 → Ln → L1 → Ln-1 → L2 → Ln-2 →

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	insertPart := head
	for i := 0; i < length/2; i++ {
		insertPart = insertPart.Next
	}
	cur := head
	var reodrerIteration func(*ListNode)
	reodrerIteration = func(node *ListNode) {
		if node != nil {
			reodrerIteration(node.Next)
		} else {
			return
		}

		if cur == node {
			return
		}
		tmp := cur.Next
		cur.Next = node
		node.Next = tmp
		cur = tmp
		if cur.Next == node {
			cur.Next = nil
		}
	}

	reodrerIteration(insertPart)
}

// func main() {
// 	list := populateList([]int{1, 2, 3, 4})
// 	reorderList(list)
// 	list.Print() // 1 4 2 3

// 	list = populateList([]int{1, 2, 3, 4, 5, 6})
// 	reorderList(list)
// 	list.Print() // 1 6 2 5 3 4

// 	list = populateList([]int{1, 2, 3, 4, 5, 6, 7})
// 	reorderList(list)
// 	list.Print() // 1 6 2 5 3 4

// 	reorderList(nil)

// 	list = populateList([]int{2, 4, 6, 8, 10})
// 	reorderList(list)
// 	list.Print() // 2 10 4 8 6
// }
