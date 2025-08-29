package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func decodeTree(args []any) *TreeNode {
	if len(args) == 0 {
		return nil
	}

	root := &TreeNode{Val: args[0].(int)}
	queue := []*TreeNode{root}

	for i := 1; i < len(args); i += 2 {
		currentElement := queue[0]
		queue = queue[1:]

		left := args[i]
		if left != nil {
			currentElement.Left = &TreeNode{Val: left.(int)}
			queue = append(queue, currentElement.Left)
		}

		if i+1 < len(args) {
			right := args[i+1]
			if right != nil {
				currentElement.Right = &TreeNode{Val: right.(int)}
				queue = append(queue, currentElement.Right)
			}
		}
	}

	return root
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
