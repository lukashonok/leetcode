package main

func kthSmallest(root *TreeNode, k int) int {
	lineTree, result := []int{}, root.Val

	var kthSmallestIteration func(*TreeNode)
	kthSmallestIteration = func(node *TreeNode) {
		if node == nil {
			return
		}

		kthSmallestIteration(node.Left)
		lineTree = append(lineTree, node.Val)
		if len(lineTree) == k {
			result = node.Val
		}
		kthSmallestIteration(node.Right)
	}

	kthSmallestIteration(root)

	return result
}

// func main() {
// 	fmt.Println(kthSmallest(decodeTree([]any{3, 1, 4, nil, 2}), 1))            // 2
// 	fmt.Println(kthSmallest(decodeTree([]any{5, 3, 6, 2, 4, nil, nil, 1}), 3)) // 3
// }
