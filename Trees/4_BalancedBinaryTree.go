package main

import "math"

func isBalanced(root *TreeNode) bool {
	isBalanced := true
	var iteration func(*TreeNode) int
	iteration = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := iteration(node.Left) + 1
		rightHeight := iteration(node.Right) + 1
		if isBalanced {
			isBalanced = math.Abs(float64(leftHeight-rightHeight)) <= 1
		}
		return max(leftHeight, rightHeight)
	}
	iteration(root)
	return isBalanced
}
