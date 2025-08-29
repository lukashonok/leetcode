package main

import (
	"math"
)

func isValidBST(root *TreeNode) bool {
	var iteration func(*TreeNode, int, int) bool
	iteration = func(node *TreeNode, minValue int, maxValue int) bool {
		if node == nil {
			return true
		}

		left := iteration(node.Left, minValue, min(node.Val, maxValue)-1)
		right := iteration(node.Right, max(node.Val, minValue)+1, maxValue)

		if node.Val < minValue || node.Val > maxValue {
			return false
		}

		return left && right
	}
	return iteration(root, -math.MaxInt, math.MaxInt)
}

// func main() {
// 	fmt.Println(isValidBST(decodeTree([]any{2, 2, 2})))                           // false
// 	fmt.Println(isValidBST(decodeTree([]any{10, 5, 20, nil, nil, 9, 25})))        // false
// 	fmt.Println(isValidBST(decodeTree([]any{2, 0, 4, nil, 1, 3, 5})))             // true
// 	fmt.Println(isValidBST(decodeTree([]any{5, 14, nil, 1})))                     // false
// 	fmt.Println(isValidBST(decodeTree([]any{5, 2, 6, 0, 3, nil, nil, nil, nil}))) // 3
// 	fmt.Println(isValidBST(decodeTree([]any{2, 1, 3, 0, 4})))                     // 3
// 	fmt.Println(isValidBST(decodeTree([]any{5, 1, 4, nil, nil, 3, 6})))           // 3
// 	fmt.Println(isValidBST(decodeTree([]any{2, 1, 3})))                           // 4

// 	fmt.Println(isValidBST(decodeTree([]any{0, nil, -1}))) // 4
// }
