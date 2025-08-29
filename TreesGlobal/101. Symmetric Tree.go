package main

func cloneTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	cloneNode := &TreeNode{Val: node.Val}
	cloneNode.Left = cloneTree(node.Left)
	cloneNode.Right = cloneTree(node.Right)
	return cloneNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSymmetric(root *TreeNode) bool {
	return isSameTree(invertTree(cloneTree(root)), root)
}

// func main() {
// 	fmt.Println(isSymmetric(decodeTree([]any{1, 2, 2, 3, 4, 4, 3})))
// }
