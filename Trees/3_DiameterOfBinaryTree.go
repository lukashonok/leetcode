package main

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	var iteration func(node *TreeNode) int
	iteration = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		maxDepthLeft := iteration(node.Left)
		maxDepthRight := iteration(node.Right)
		maxDepth := max(maxDepthLeft, maxDepthRight)
		diameter = max(maxDepthLeft+maxDepthRight, diameter)
		return maxDepth + 1
	}
	iteration(root)
	return diameter
}
