package main

type Pair struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*Pair{{node: root, level: 0}}

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		if len(result) < pair.level+1 {
			result = append(result, []int{})
		}

		result[pair.level] = append(result[pair.level], pair.node.Val)
		if pair.node.Left != nil {
			queue = append(queue, &Pair{node: pair.node.Left, level: pair.level + 1})
		}
		if pair.node.Right != nil {
			queue = append(queue, &Pair{node: pair.node.Right, level: pair.level + 1})
		}
	}

	return result
}

// func main() {
// 	fmt.Println(levelOrder(decodeTree([]any{3, 9, 20, nil, nil, 15, 7}))) // [[3],[9,20],[15,7]]
// 	fmt.Println(levelOrder(decodeTree([]any{1, 2, 3, 4, 5, 6, 7})))       // [[1],[2,3],[4,5,6,7]]
// 	fmt.Println(levelOrder(decodeTree([]any{1})))                         // [[1]]
// 	fmt.Println(levelOrder(decodeTree([]any{})))                          // []
// }
