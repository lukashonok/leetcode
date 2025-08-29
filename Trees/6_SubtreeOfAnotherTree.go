package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currentElement := queue[0]
		queue = queue[1:]
		if currentElement.Val == subRoot.Val && isSameTree(currentElement, subRoot) {
			return true
		}
		if currentElement.Left != nil {
			queue = append(queue, currentElement.Left)
		}
		if currentElement.Right != nil {
			queue = append(queue, currentElement.Right)
		}
	}

	return false
}
