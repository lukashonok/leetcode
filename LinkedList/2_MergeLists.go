package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var merged *ListNode
	currentMerged, list1Node, list2Node := merged, list1, list2
	for list1Node != nil || list2Node != nil {
		newNode := &ListNode{}

		if list1Node == nil || (list2Node != nil && list2Node.Val < list1Node.Val) {
			newNode.Val = list2Node.Val
			list2Node = list2Node.Next
		} else {
			newNode.Val = list1Node.Val
			list1Node = list1Node.Next
		}

		if currentMerged == nil {
			currentMerged = &ListNode{Val: newNode.Val}
			merged = currentMerged
		} else {
			currentMerged.Next = newNode
			currentMerged = currentMerged.Next
		}
	}

	return merged
}

// func main() {
// 	mergeTwoLists(populateList([]int{1, 2, 4}), populateList([]int{1, 3, 4})).Print()   // 1 1 2 3 4 4
// 	mergeTwoLists(populateList([]int{10, 20}), populateList([]int{30, 40, 50})).Print() // 10 20 30 40 50
// 	mergeTwoLists(populateList([]int{30, 40, 50}), populateList([]int{10, 20})).Print() // 10 20 30 40 50
// 	mergeTwoLists(nil, nil).Print()                                                     //
// 	mergeTwoLists(nil, populateList([]int{0})).Print()                                  // 0
// 	mergeTwoLists(nil, populateList([]int{1, 2})).Print()                               // 1 2
// }
