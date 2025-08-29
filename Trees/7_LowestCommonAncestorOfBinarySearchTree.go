package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for {
		if root.Val > p.Val && root.Val > q.Val {
			// To the left both
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			// To the right both
			root = root.Right
		} else {
			return root
		}
	}
}

// func main() {
// 	root := decodeTree([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5})
// 	q := root.Left.Left
// 	p := root.Left.Right.Right
// 	fmt.Println(lowestCommonAncestor(root, p, q))
// 	p = root.Left
// 	q = root.Left.Right
// 	fmt.Println(lowestCommonAncestor(root, p, q))
// 	root = decodeTree([]any{2, 1})
// 	p = root
// 	q = root.Left
// 	fmt.Println(lowestCommonAncestor(root, p, q))
// }
