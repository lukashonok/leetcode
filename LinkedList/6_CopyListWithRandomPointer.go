package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (head *Node) Print() {
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

func populate(input [][]int) *Node {
	if len(input) == 0 {
		return nil
	}
	head := &Node{Val: input[0][0]}
	cur := head
	for i := 1; i < len(input); i++ {
		next := &Node{Val: input[i][0]}
		cur.Next = next
		cur = next
	}
	globalCur := head
	for _, v := range input {
		if len(v) == 2 {
			cur := head
			for i := 0; i < v[1]; i++ {
				cur = cur.Next
			}
			globalCur.Random = cur
		}
		globalCur = globalCur.Next
	}
	return head
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	randomElementsMap := map[*Node]*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		if _, alreadyExists := randomElementsMap[cur.Random]; cur.Random != nil && !alreadyExists {
			randomElementsMap[cur.Random] = &Node{Val: cur.Random.Val}
		}
	}

	var newHead *Node
	copyCur := newHead
	lastCreated := copyCur
	for cur := head; cur != nil; cur = cur.Next {
		if _, alreadyExists := randomElementsMap[cur]; !alreadyExists {
			copyCur = &Node{Val: cur.Val}
		} else {
			copyCur = randomElementsMap[cur]
		}

		if cur.Random != nil {
			copyCur.Random = randomElementsMap[cur.Random]
		}

		if lastCreated != nil {
			lastCreated.Next = copyCur
		} else {
			newHead = copyCur
		}

		lastCreated = copyCur
	}

	return newHead
}

// func main() {
// 	copyRandomList(populate([][]int{{7}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})).Print()
// 	copyRandomList(populate([][]int{})).Print()
// 	copyRandomList(populate([][]int{{1, 1}, {2, 1}})).Print()
// 	copyRandomList(populate([][]int{{3}, {3, 0}, {3}})).Print()
// }
