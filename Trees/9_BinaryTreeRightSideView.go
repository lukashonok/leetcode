package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		result = append(result, queue[size-1].Val)
		for range size {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

// func main() {
// 	fmt.Println(rightSideView(decodeTree([]any{1, 2, 3, nil, 5, nil, 4})))      // [1,3,4]
// 	fmt.Println(rightSideView(decodeTree([]any{1, 2, 3, 4, nil, nil, nil, 5}))) // [1,3,4,5]
// 	fmt.Println(rightSideView(decodeTree([]any{1, nil, 3})))                    // [1,3]
// 	fmt.Println(rightSideView(decodeTree([]any{})))                             // []
// }
