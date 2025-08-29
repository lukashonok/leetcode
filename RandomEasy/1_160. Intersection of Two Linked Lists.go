package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	passed := map[*ListNode]bool{}
	for headA != nil || headB != nil {
		if headA != nil {
			_, exists := passed[headA]
			if exists {
				return headA
			}
			passed[headA] = true
			headA = headA.Next
		}
		if headB != nil {
			_, exists := passed[headB]
			if exists {
				return headB
			}
			passed[headB] = true
			headB = headB.Next
		}
	}
	return nil
}

// func main() {

// }
