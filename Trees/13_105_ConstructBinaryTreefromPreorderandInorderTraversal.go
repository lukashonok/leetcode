package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap, preorderNodes, length := map[int]int{}, []*TreeNode{}, len(preorder)
	for i, v := range inorder {
		inorderMap[v] = i
		preorderNodes = append(preorderNodes, &TreeNode{Val: preorder[i]})
	}

	preorderIndex := 0
	var iteration func(*TreeNode, int, int) *TreeNode
	iteration = func(node *TreeNode, l int, r int) *TreeNode {
		if l == r {
			return node
		}

		inorderIndex := inorderMap[node.Val]
		if preorderIndex+1 < length {
			preorderIndex++
			node.Left = iteration(preorderNodes[preorderIndex], l, inorderIndex-1)
		}
		if preorderIndex+1 < length {
			preorderIndex++
			node.Right = iteration(preorderNodes[preorderIndex], inorderIndex+1, r)
		}

		return node
	}

	a := iteration(preorderNodes[0], 0, len(inorder)-1)
	return a
}

// func main() {
// 	fmt.Println(buildTree([]int{3, 2, 1, 4}, []int{1, 2, 3, 4}))
// 	fmt.Println(buildTree([]int{1, 2}, []int{1, 2}))
// 	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
// 	// fmt.Println(buildTree([]int{1, 2}, []int{2, 1}))
// 	fmt.Println(buildTree([]int{2, 3, 5, 16, 4, 7, 8}, []int{16, 3, 4, 2, 5, 8, 7}))
// 	fmt.Println(buildTree([]int{-1}, []int{-1}))
// }
