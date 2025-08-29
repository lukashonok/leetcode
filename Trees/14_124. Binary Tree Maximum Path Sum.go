package main

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := root.Val

	var iteration func(*TreeNode) int
	iteration = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSide := iteration(node.Left)
		rightSide := iteration(node.Right)
		possiblePathSum := max(leftSide+node.Val+rightSide, leftSide+node.Val, node.Val+rightSide, node.Val)

		if possiblePathSum > maxSum {
			maxSum = possiblePathSum
		}

		return max(node.Val+leftSide, node.Val+rightSide, node.Val)
	}
	iteration(root)

	return maxSum
}

// func main() {
// 	fmt.Println(maxPathSum(decodeTree([]any{9, 6, -3, nil, nil, -6, 2, nil, nil, 2, nil, -6, -6, -6}))) // 16
// 	fmt.Println(maxPathSum(decodeTree([]any{1, -2, 3})))                                                // 4
// 	fmt.Println(maxPathSum(decodeTree([]any{-10, 9, 20, nil, nil, 15, 7})))                             // 42
// 	fmt.Println(maxPathSum(decodeTree([]any{1, 2, 3})))                                                 // 6
// }
