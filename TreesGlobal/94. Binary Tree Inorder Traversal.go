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

func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	var iteration func(node *TreeNode)
	iteration = func(node *TreeNode) {
		if node == nil {
			return
		}
		iteration(node.Left)
		result = append(result, node.Val)
		iteration(node.Right)
	}
	iteration(root)

	return result
}

// func main() {
// 	fmt.Println(inorderTraversal(decodeTree([]any{1, nil, 2, 3})))
// 	fmt.Println(inorderTraversal(decodeTree([]any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9})))
// }
