package main

func hasCycle(head *ListNode) bool {
	for tort, hare := head, head.Next; tort != nil && hare != nil && hare.Next != nil; tort, hare = tort.Next, hare.Next.Next {
		if tort == hare {
			return true
		}
	}
	return false
}

// func main() {
// 	// 3,2,0,-4
// 	a4 := &ListNode{Val: -4}
// 	a3 := &ListNode{Val: 0, Next: a4}
// 	a2 := &ListNode{Val: 2, Next: a3}
// 	a1 := &ListNode{Val: 3, Next: a2}
// 	a4.Next = a2
// 	fmt.Println(hasCycle(a1))
// }
