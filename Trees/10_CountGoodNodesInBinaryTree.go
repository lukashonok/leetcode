package main

func goodNodes(root *TreeNode) int {
	total := 0
	var iteration func(*TreeNode, int)
	iteration = func(node *TreeNode, current int) {
		if node == nil {
			return
		}

		if node.Val >= current {
			total++
		}

		iteration(node.Left, max(node.Val, current))
		iteration(node.Right, max(node.Val, current))
	}
	iteration(root, root.Val)
	return total
}

// func main() {
// 	fmt.Println(goodNodes(decodeTree([]any{3, 1, 4, 3, nil, 1, 5}))) // 4
// 	fmt.Println(goodNodes(decodeTree([]any{3, 3, nil, 4, 2})))       // 3
// }
