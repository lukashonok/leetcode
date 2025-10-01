package main

import "fmt"

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	if root == nil {
		return newNode
	}
	var iteration func(node *TreeNode)
	iteration = func(node *TreeNode) {
		if val < node.Val {
			if node.Left != nil {
				iteration(node.Left)
			} else {
				node.Left = newNode
			}
		} else {
			if node.Right != nil {
				iteration(node.Right)
			} else {
				node.Right = newNode
			}
		}

	}
	iteration(root)
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			root.Val = cur.Val
			root.Right = deleteNode(root.Right, cur.Val)
		}
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var iteration func(node *TreeNode) (int, int)
	iteration = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		withRootLeft, withoutRootLeft := iteration(node.Left)
		withRootRight, withoutRootRight := iteration(node.Right)

		withRoot := node.Val + withoutRootLeft + withoutRootRight
		withoutRoot := max(withRootLeft, withoutRootLeft) + max(withRootRight, withoutRootRight)
		return withRoot, withoutRoot
	}
	return max(iteration(root))
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var iteration func(node *TreeNode) (*TreeNode, bool)
	iteration = func(node *TreeNode) (*TreeNode, bool) {
		if node == nil {
			return nil, true
		}
		left, leftLeaf := iteration(node.Left)
		right, rightLeaf := iteration(node.Right)
		if left != nil && left.Val == target && leftLeaf {
			node.Left = nil
		}
		if right != nil && right.Val == target && rightLeaf {
			node.Right = nil
		}
		return node, node.Left == nil && node.Right == nil
	}
	root, rootLeaf := iteration(root)
	if root != nil && root.Val == target && rootLeaf {
		return nil
	}
	return root
}

func main() {
	fmt.Println(removeLeafNodes(decodeTree([]any{1, 2, 3, 2, nil, 2, 4}), 2))
	// fmt.Println(rob(decodeTree([]any{3, 2, 3, nil, 3, nil, 1})))
}
