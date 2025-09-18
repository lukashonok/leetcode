package main

import (
	"fmt"
	"slices"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	divisionPoint := slices.Index(inorder, node.Val)
	node.Left = buildTree(preorder[1:divisionPoint+1], inorder[:divisionPoint])
	node.Right = buildTree(preorder[divisionPoint+1:], inorder[divisionPoint+1:])
	return node
}

func main() {
	// fmt.Println(buildTree([]int{3, 2, 1, 4}, []int{1, 2, 3, 4}))
	// fmt.Println(buildTree([]int{1, 2}, []int{1, 2}))
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	// fmt.Println(buildTree([]int{1, 2}, []int{2, 1}))
	// fmt.Println(buildTree([]int{2, 3, 5, 16, 4, 7, 8}, []int{16, 3, 4, 2, 5, 8, 7}))
	// fmt.Println(buildTree([]int{-1}, []int{-1}))
}
